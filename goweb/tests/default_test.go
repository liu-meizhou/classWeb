package tests

import (
	"fmt"
	"github.com/extrame/xls"
	"testing"
)

// TestReadXls 解析xls放入数据库
func TestReadXls(t *testing.T) {
	if xlFile, err := xls.Open("班级课表/计科182课表1.xls", "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				for j := 0; j <= (row1.LastCol()); j++ {
					cell := row1.Col(j)
					fmt.Print("\n", "第 ", i, " 行,第 ", j, " 列: 数据: ", cell)
				}
			}
		}
	}
}

