package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// 不安全协议颜色
func insecureProtocolColor(s string) tablewriter.Colors {
	if s == "Yes" {
		return tablewriter.Colors{tablewriter.FgRedColor}
	} else if s == "No" {
		return tablewriter.Colors{tablewriter.FgGreenColor}
	} else {
		return tablewriter.Colors{tablewriter.FgYellowColor}
	}
}

// 安全协议颜色
func secureProtocolColor(s string) tablewriter.Colors {
	if s == "Yes" {
		return tablewriter.Colors{tablewriter.FgGreenColor}
	} else if s == "No" {
		return tablewriter.Colors{tablewriter.FgRedColor}
	} else {
		return tablewriter.Colors{tablewriter.FgYellowColor}
	}
}

// 结果颜色
func resultColor(s string) tablewriter.Colors {
	if s == "PASS" {
		return tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor}
	} else if s == "FAIL" {
		return tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor}
	} else {
		return tablewriter.Colors{tablewriter.Bold, tablewriter.BgYellowColor}
	}
}

// 渲染表格
func tableRender(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No.", "Site", "IP", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3", "Result"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	passCount := 0
	failCount := 0
	unknownCount := 0
	for i, row := range data {
		row = append([]string{strconv.Itoa(i + 1)}, row...)
		if row != nil {
			switch row[8] {
			case "PASS":
				passCount++
			case "FAIL":
				failCount++
			default:
				unknownCount++
			}
			table.Rich(row, []tablewriter.Colors{
				{},
				{},
				{},
				insecureProtocolColor(row[3]),
				insecureProtocolColor(row[4]),
				insecureProtocolColor(row[5]),
				secureProtocolColor(row[6]),
				secureProtocolColor(row[7]),
				resultColor(row[8]),
			},
			)
		}
		i++
	}
	println("\n")
	println(
		color.New(color.BgBlue, color.Bold).Sprint(" TOTAL: "+strconv.Itoa(passCount+failCount+unknownCount)+" "),
		color.New(color.BgGreen, color.Bold).Sprint(" PASS: "+strconv.Itoa(passCount)+" "),
		color.New(color.BgRed, color.Bold).Sprint(" FAIL: "+strconv.Itoa(failCount)+" "),
		color.New(color.BgYellow, color.Bold).Sprint(" UNKNOWN: "+strconv.Itoa(unknownCount)+" "),
	)
	table.Render()
}

// 渲染CSV
func csvRender(data [][]string) {
	f, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		println("文件创建失败")
		return
	}
	defer f.Close()

	f.WriteString("No.,Site,IP,SSLv3,TLSv1.0,TLSv1.1,TLSv1.2,TLSv1.3,Result\n")
	for i, row := range data {
		row = append([]string{strconv.Itoa(i + 1)}, row...)
		if row != nil {
			f.WriteString(strings.Join(row, ",") + "\n")
		}
		i++
	}
	println("\nCheck result has been saved to result.csv")
}
