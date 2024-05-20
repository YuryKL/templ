// Code generated by templ - DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Message string `json:"msg"`
}

func JSON(v any) (string, error) {
	s, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func JSONScript(id string, data any) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return err
		}
		if _, err = io.WriteString(w, `<script`); err != nil {
			return err
		}
		if id != "" {
			if _, err = fmt.Fprintf(w, ` id="%s"`, templ.EscapeString(id)); err != nil {
				return err
			}
		}
		if _, err = fmt.Fprintf(w, ` type="application/json">%s</script>`, string(dataJSON)); err != nil {
			return err
		}
		return nil
	})
}

func Page(attributeData Data, scriptData Data) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html><head><title>Script usage</title><script src=\"/assets/js/index.js\" defer></script></head><body><button id=\"attributeAlerter\" alert-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(JSON(attributeData))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `examples/typescript/components/index.templ`, Line: 49, Col: 65}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Show alert from data in alert-data attribute</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = JSONScript("scriptData", scriptData).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"scriptAlerter\">Show alert from data in script</button></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}