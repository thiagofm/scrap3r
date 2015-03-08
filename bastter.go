package main

import (
  "log"
  "github.com/PuerkitoBio/goquery"
  "github.com/gin-gonic/gin"
)

type pairData struct {
  Key string
  Value string
}

type yearlyData struct {
  Data [11]pairData
}

func BastterScrap() []yearlyData {
  doc, err := goquery.NewDocument("http://bastter.com/mercado/acao/BRFS.aspx")

  if err != nil {
    log.Fatal(err)
  }

  var arrayOfYearlyData []yearlyData

  d := [11] pairData {};
  doc.Find(".evanual").Each(func(i int, evolucaoAnual *goquery.Selection) {
    if i == 0 {
      evolucaoAnual.Find("thead tr td").Each(func(j int, tabela *goquery.Selection) {
        columnName := tabela.Text()
        d[j] = pairData{Key: columnName}
      });
      yd := yearlyData{Data: d}

      rows := evolucaoAnual.Find("tbody tr").Length()
      arrayOfYearlyData = make([]yearlyData, rows)

      evolucaoAnual.Find("tbody tr").Each(func(j int, tabela *goquery.Selection) {
        newYd := yd

        tabela.Find("td").Each(func(k int, cell *goquery.Selection) {
          newYd.Data[k].Value = cell.Text();
        })

        arrayOfYearlyData[j] = newYd
      });

    }
  })

  return arrayOfYearlyData
}

func main() {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
      c.JSON(200, BastterScrap())
  })
  router.Run(":8080")
}
