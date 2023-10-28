// Code generated by templ@v0.2.408 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Text(letters []string, author string) templ.Component {
	return templ.ComponentFunc(func(templ_7745c5c3_Ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		templ_7745c5c3_Ctx = templ.InitializeContext(templ_7745c5c3_Ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(templ_7745c5c3_Ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		templ_7745c5c3_Ctx = templ.ClearChildren(templ_7745c5c3_Ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := `
    .correct {
      color: green;
    }

    .incorrect {
      color: red;
    }

    .incorrect-space {
      background-color: red;
    }

    #active {
      text-decoration: underline;
    }
  `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</style><script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := `
  function next() {
    htmx.trigger("#text-container", "manual")
  }
  `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</script><div class=\"text\"><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `WPM: `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span id=\"wpm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var5 := `0`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></p><p id=\"text-paragraph\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i, letter := range letters {
			if i == 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"letter\" id=\"active\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string = letter
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"letter\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string = letter
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := `- `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string = author
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(templ_7745c5c3_Ctx, templ_7745c5c3_Buffer, interact())
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div on=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 templ.ComponentScript = interact()
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var10.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = onLoad(interact()).Render(templ_7745c5c3_Ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func interact() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_interact_f2a4`,
		Function: `function __templ_interact_f2a4(){const textParagraph = document.getElementById("text-paragraph");
  const letters = textParagraph.getElementsByClassName("letter");

  let hasStarted = false;
  let startTime = undefined;
  let wordsTyped = 0;
  let isAllCorrect = false;
  let intervalID = undefined;

  addEventListener("keydown", onChange);

  function startInterval() {
    return setInterval(() => {
      if (hasStarted) {
        const time = Date.now() - startTime;
        const wpm = Math.round(wordsTyped / (time / 1000 / 60));
        document.getElementById("wpm").innerText = wpm;
      }
    }, 500);
  }

  function onChange(e) {
    if (event.key === "Enter" && isAllCorrect) {
      next()
    }
    if ((event.key.length != 1 && event.key != "Backspace") || isAllCorrect) {
      return
    }

    // Get index of active letter
    let i = 0;
    for (const letter of letters) {
      if (letter.id === "active") {
        break;
      }
      i++
    }

    if (e.key === "Backspace") {
      if (i > 0 ) {
        letters[i - 1].classList.remove("correct");
        letters[i - 1].classList.remove("incorrect");
        letters[i - 1].classList.remove("incorrect-space");
        letters[i].id = "";
        letters[i - 1].id = "active";
        if (letters[i].innerText === " ") {
          wordsTyped--;
        }
      }
      if (i === 1 ) {
        clearInterval(intervalID);
        hasStarted = false;
        startTime = undefined;
        wordsTyped = 0;
        document.getElementById("wpm").innerText = "0";
      }
      return
    }

    if (!hasStarted) {
      hasStarted = true;
      startTime = Date.now();
      intervalID = startInterval();
    }

    // Handle current letter
    const currentLetter = letters[i];
    if (e.key === currentLetter.innerText) {
      currentLetter.classList.add("correct");
      if (i < letters.length - 1) {
        currentLetter.id = "";
        letters[i + 1].id = "active";
      }

      if (e.key === " ") {
        wordsTyped++;
      }
    } else {
      currentLetter.classList.add("incorrect");
      if (currentLetter.innerText === " ") {
      currentLetter.classList.add("incorrect-space");  
      }
      if (i < letters.length - 1) {
        currentLetter.id = "";
        letters[i + 1].id = "active";
      }
      if (e.key === " ") {
        wordsTyped--;
      }
      return
    }

    // If on last letter, check if game is won
    if (i === letters.length - 1) {
      let checkAllCorrect = true;
      for (const l of letters) {
        if (!l.classList.contains("correct")) {
          checkAllCorrect = false;
          break;
        }
      }

      if (checkAllCorrect) {
        isAllCorrect = true;
        clearInterval(intervalID);
      }
    }
  }}`,
		Call: templ.SafeScript(`__templ_interact_f2a4`),
	}
}

func onLoad(script templ.ComponentScript) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		if _, err = io.WriteString(w, `<script type="text/javascript">`+"\r\n"+script.Call+"\r\n</script>"); err != nil {
			return err
		}
		return nil
	})
}
