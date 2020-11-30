package main

// Don't go snooping around here for answers! Puzzle spoilers lie below.

import (
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.NoRoute(func(c *gin.Context) {
        relativePath := string(c.Request.URL.String())
        pathElems := strings.Split(relativePath, "/")

        if pathElems[1] == "puzzle" {
            if len(pathElems) >= 3 {
                // Serve the "wrong answer" page with a hint (if applicable).
                c.HTML(http.StatusOK, "wrong-answer.tmpl", gin.H{
                    "wronganswer": pathElems[2],
                    "hint": getHint(pathElems[2]),
                })
            } else {
                // Serve the "wrong answer" page with the empty string.
                c.HTML(http.StatusOK, "wrong-answer.tmpl", gin.H{
                    "wronganswer": "",
                })
            }
        }

        // Serve a normal 404 page.
        // TODO: Make a not-ugly 404 page.
    })

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Page 1",
        })
    })

    router.GET("/puzzle/1", func(c *gin.Context) {
        c.HTML(http.StatusOK, "1.tmpl", gin.H{})
    })

    router.GET("/puzzle/2", func(c *gin.Context) {
        c.HTML(http.StatusOK, "2.tmpl", gin.H{})
    })

    router.GET("/puzzle/taste", func(c *gin.Context) {
        c.HTML(http.StatusOK, "taste.tmpl", gin.H{})
    })

    router.GET("/puzzle/smartphone", func(c *gin.Context) {
        c.HTML(http.StatusOK, "smartphone.tmpl", gin.H{})
    })

    router.GET("/puzzle/gifts", func(c *gin.Context) {
        c.HTML(http.StatusOK, "gifts.tmpl", gin.H{})
    })

    port := os.Getenv("PORT") // Define this on Heroku.
    router.Run(":" + port)
}


// Text to add to a 404 page when the URL is equal to /puzzle/${DICT_KEY}.
var hints = map[string]string{
    "3":            "You didn't think it would be that easy, did you?",
    "after":        "After what?",
    "aftertaste":   "You're onto something.",
    "phone":        "What would you call a smart one of these? (You pretty much solved it)",
}
func getHint(wronganswer string) string {
    if val, ok := hints[wronganswer]; ok {
        return "** " + val + " **"
    }
    return ""
}


/* TABLE OF CONTENTS
   (so we remember which order the puzzles come in)

   PUZZLE #, FILENAME
   1            1
   2            2
   3            taste
   taste        smartphone
   smartphone   gifts
   gifts        ...
*/
