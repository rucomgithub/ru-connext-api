package studenth
import (
	"time"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandlers) GetPhoto(c *gin.Context) {

    timeout := time.Duration(5*time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    surl := "http://10.3.68.55/ita/images/Logo_Color_600.png"

    req, err := http.NewRequest("GET", surl, nil)
    // AccessToken := c.Request().Header.Get("Authorization")
    // fmt.Println(AccessToken)
    // req.Header.Set("Authorization", AccessToken)

    if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"Get Photo"})
		c.Abort()
		return
    }
    resp, err := client.Do(req)
    if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"client.Do"})
		c.Abort()
		return
    }

    f, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message":"ioutil.ReadAll"})
		c.Abort()
		return
    }

    resp.Body.Close()

    c.Data(http.StatusOK, "image/jpeg", f)
    return 
}