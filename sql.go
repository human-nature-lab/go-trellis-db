package main

import (
	"fmt"
	"io"
	"strings"
)

func writeEdgeInsert(out io.Writer, chunk []Edge) (err error) {
	cols := []string{"id", "source_respondent_id", "target_respondent_id", "created_at", "updated_at", "note"}
	if len(chunk) == 0 {
		return
	}
	if _, err = fmt.Fprintf(out, "INSERT INTO edge (%s) VALUES \n", strings.Join(cols, ", ")); err != nil {
		return
	}
	for i, r := range chunk {
		if r.Note.Valid {
			if _, err = fmt.Fprintf(out, "(\"%s\", \"%s\", \"%s\", now(), now(), \"%s\")", r.Id, r.SourceRespondentId, r.TargetRespondentId, r.Note.String); err != nil {
				return
			}
		} else {
			if _, err = fmt.Fprintf(out, "(\"%s\", \"%s\", \"%s\", now(), now(), null)", r.Id, r.SourceRespondentId, r.TargetRespondentId); err != nil {
				return
			}
		}
		if i != len(chunk)-1 {
			if _, err = fmt.Fprintln(out, ","); err != nil {
				return
			}
		}
	}
	_, err = fmt.Fprintln(out, ";\n")
	return
}

func writePreloadInsert(out io.Writer, chunk []PreloadAction, questionId string) (err error) {
	cols := []string{"id", "respondent_id", "question_id", "payload", "action_type", "created_at"}
	if len(chunk) == 0 {
		return
	}
	if _, err = fmt.Fprintf(out, "INSERT INTO preload_action (%s) VALUES \n", strings.Join(cols, ", ")); err != nil {
		return
	}
	for i, r := range chunk {
		if _, err = fmt.Fprintf(out, "(\"%s\", \"%s\", \"%s\", \"%s\", \"add-edge\", now())", r.Id, r.RespondentId, questionId, strings.ReplaceAll(r.Payload, `"`, `\"`)); err != nil {
			return
		}
		if i != len(chunk)-1 {
			if _, err = fmt.Fprintln(out, ","); err != nil {
				return
			}
		}
	}
	_, err = fmt.Fprintln(out, ";\n")
	return
}

func WriteEdges(out io.Writer, edgeMap map[string]Edge) (err error) {
	chunk := []Edge{}
	for _, edge := range edgeMap {
		chunk = append(chunk, edge)
		if len(chunk) >= 2000 {
			if err = writeEdgeInsert(out, chunk); err != nil {
				return
			}
			chunk = chunk[:0]
		}
	}
	return writeEdgeInsert(out, chunk)
}

func WritePreloads(out io.Writer, preloads []PreloadAction, questionId string) (err error) {
	chunk := []PreloadAction{}
	for _, edge := range preloads {
		chunk = append(chunk, edge)
		if len(chunk) >= 2000 {
			if err = writePreloadInsert(out, chunk, questionId); err != nil {
				return
			}
			chunk = chunk[:0]
		}
	}
	return writePreloadInsert(out, chunk, questionId)
}
