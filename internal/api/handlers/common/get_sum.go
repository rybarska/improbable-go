package common

import ( "fmt" "net/http"

"allaboutapps.dev/aw/go-starter/internal/api"
"github.com/labstack/echo/v4"
)

func GetSumRoute(s api.Server) echo.Route { return s.Router.Management.GET("/sum/:count", getSumHandler(s)) }

// :count "1" => 1 => "1" // :count "2" => 1+2 => "3" // :count "3" => 1+2+3 => "6" func getSumHandler(s *api.Server) echo.HandlerFunc { return func(c echo.Context) error {

    var cnt int64

    if err := echo.PathParamsBinder(c).Int64("count", &cnt).BindError(); err != nil {
        return c.String(http.StatusBadRequest, "Please provide an integer.\n")
    }

    if cnt == 0 {
        return c.String(http.StatusOK, "0\n")
    }

    var sum int64 = 1

    var i int64
    for i = 1; i != cnt; i++ {
        sum += (i + 1)
    }

    return c.String(http.StatusOK, fmt.Sprint(sum, "\n"))
