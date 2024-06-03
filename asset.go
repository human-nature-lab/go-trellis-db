package go_trellis_db

import (
	"crypto/md5"
	"io"
	"mime"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/wyattis/z/zset/zstringset"
)

type AssetType string

const (
	AssetTypeImage AssetType = "image"
	AssetTypeVideo AssetType = "video"
	AssetTypeAudio AssetType = "audio"
	AssetTypeText  AssetType = "text"
	AssetTypeOther AssetType = "other"
)

var imageTypes = zstringset.New("image/jpeg", "image/png", "image/gif", "image/svg+xml", "image/webp")
var videoTypes = zstringset.New("video/mp4", "video/webm", "video/ogg")
var audioTypes = zstringset.New("audio/mpeg", "audio/ogg", "audio/wav", "audio/webm", "audio/aac", "audio/flac")
var textTypes = zstringset.New("text/plain", "text/html", "text/css", "text/javascript", "application/json",
	"application/xml", "application/xhtml+xml", "application/rss+xml", "application/atom+xml", "application/ld+json",
	"application/manifest+json", "application/javascript", "application/x-javascript")

func MimeToType(mimeType string) AssetType {
	if imageTypes.Contains(mimeType) {
		return AssetTypeImage
	}
	if videoTypes.Contains(mimeType) {
		return AssetTypeVideo
	}
	if audioTypes.Contains(mimeType) {
		return AssetTypeAudio
	}
	if textTypes.Contains(mimeType) {
		return AssetTypeText
	}
	return AssetTypeOther
}

func CreateAssetFromPath(assetRoot string, db *sqlx.DB, path string) (asset Asset, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return
	}
	return CreateAsset(assetRoot, db, file, filepath.Base(path), stat.Size())
}

func CreateAsset(assetRoot string, db *sqlx.DB, file io.Reader, fileName string, size int64) (asset Asset, err error) {
	asset.Id = uuid.New().String()
	asset.FileName = fileName
	asset.MimeType = mime.TypeByExtension(filepath.Ext(fileName))
	asset.Type = string(MimeToType(asset.MimeType))
	asset.Size = int(size)
	asset.CreatedAt = LaravelNow()
	asset.UpdatedAt = LaravelNow()

	sum := md5.New()
	reader := io.TeeReader(file, sum)

	assetPath := filepath.Join(assetRoot, asset.Id)
	f, err := os.Create(assetPath)
	if err != nil {
		return
	}
	defer f.Close()
	if _, err = io.Copy(f, reader); err != nil {
		return
	}
	asset.Md5Hash = string(sum.Sum(nil))
	_, err = db.NamedExec(`INSERT INTO asset (id, file_name, mime_type, type, size, md5_hash) VALUES (:id, :file_name, :mime_type, :type, :size, :md5_hash)`, asset)
	return
}
