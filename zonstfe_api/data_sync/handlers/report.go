package handlers

import (
	"net/http"
	"zonstfe_api/common/my_context"
	"zonstfe_api/common/utils/myfile"
	"path/filepath"
	"strconv"
	"time"
	"io"
	"os"
	"fmt"
	"strings"
	"github.com/satori/go.uuid"
	"github.com/pkg/errors"
	"path"
	"zonstfe_api/corm"
	"reflect"
)

const root_path = "/tmp/"

func (c *Handler) ReportAppSlotImport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Url *string `form:"url"`
	}
	if err := c.BindQuery(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Url == nil {
		c.JsonError(w, my_context.ErrBadValid, nil)
		return
	}
	event_id := uuid.NewV4().String()
	// 事物开始
	c.LogEventStart("app广告位报表导入", event_id, *req.Url, time.Now().Unix())

	resp, err := http.Get(*req.Url)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	defer resp.Body.Close()
	// 存储文件
	file_ext := filepath.Ext(*req.Url)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + path.Base(*req.Url)
	file, err := os.OpenFile(root_path+filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	table_name := "report_app_slot"
	columns := []string{"os", "bundle_id", "app_key",
		"slot_id", "report_date", "imp", "clk"}
	if err := c.CopyReportToDb(table_name, file_ext, root_path+filename, columns); err != nil {
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}
	c.LogEventEnd(event_id, "", 1, time.Now().Unix())
	c.JsonBase(w, nil)

}

func (c *Handler) ReportAppRewardImport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Url *string `form:"url"`
	}
	if err := c.BindQuery(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Url == nil {
		c.JsonError(w, my_context.ErrBadValid, nil)
		return
	}
	event_id := uuid.NewV4().String()
	// 事物开始
	c.LogEventStart("app激励报表导入", event_id, *req.Url, time.Now().Unix())

	resp, err := http.Get(*req.Url)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	defer resp.Body.Close()
	// 存储文件
	file_ext := filepath.Ext(*req.Url)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + path.Base(*req.Url)
	file, err := os.OpenFile(root_path+filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	table_name := "report_app_reward"
	columns := []string{"os", "bundle_id", "app_key",
		"report_date", "imp", "amount", "reward", "uv"}
	if err := c.CopyReportToDb(table_name, file_ext, root_path+filename, columns); err != nil {
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}
	c.LogEventEnd(event_id, "", 1, time.Now().Unix())
	c.JsonBase(w, nil)

}

func (c *Handler) ReportAppImport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Url *string `form:"url"`
	}
	if err := c.BindQuery(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Url == nil {
		c.JsonError(w, my_context.ErrBadValid, nil)
		return
	}
	event_id := uuid.NewV4().String()
	// 事物开始
	c.LogEventStart("app报表导入", event_id, *req.Url, time.Now().Unix())
	fmt.Println(*req.Url)
	resp, err := http.Get(*req.Url)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return

	}
	defer resp.Body.Close()
	// 存储文件
	file_ext := filepath.Ext(*req.Url)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + path.Base(*req.Url)
	file, err := os.OpenFile(root_path+filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	lines, err := myfile.ReadLines(root_path + filename)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	args := make([][]interface{}, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		cols := make([]string, 0)
		switch file_ext {
		case ".txt":
			cols = strings.Split(line, "\t")
		case ".csv":
			cols = strings.Split(line, ",")
		}
		if len(cols) != 12 {
			c.LogEventEnd(event_id, "数据格式长度错误", -1, time.Now().Unix())
			return
		}
		cost, err := strconv.Atoi(cols[11])
		if err != nil {
			c.Logger.Println(err)
			c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
			return
		}
		cols[11] = fmt.Sprintf("%v", float64(cost)/1000000)
		args = append(args, InterfaceSlice(cols))

	}
	table_name := "report_app"
	columns := []string{"vendor_id", "user_id", "campaign_id",
		"ad_id", "os", "bundle_id", "report_date", "win", "imp", "clk", "eimp", "cost"}
	//if err := c.CopyReportToDb(table_name, file_ext, root_path+filename, columns); err != nil {
	//	c.Logger.Println(err)
	//	c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
	//	return
	//}
	if err := corm.PgCopy(table_name, args, columns); err != nil {
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}
	c.LogEventEnd(event_id, "", 1, time.Now().Unix())
	c.JsonBase(w, nil)

}
func (c *Handler) ReportBaseImport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Url *string `form:"url"`
	}
	if err := c.BindQuery(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Url == nil {
		c.JsonError(w, my_context.ErrBadValid, nil)
		return
	}
	event_id := uuid.NewV4().String()
	// 事物开始
	c.LogEventStart("基本报表导入", event_id, *req.Url, time.Now().Unix())

	resp, err := http.Get(*req.Url)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	defer resp.Body.Close()
	// 存储文件
	file_ext := filepath.Ext(*req.Url)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + path.Base(*req.Url)
	file, err := os.OpenFile(root_path+filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	lines, err := myfile.ReadLines(root_path + filename)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	args := make([][]interface{}, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		cols := make([]string, 0)
		switch file_ext {
		case ".txt":
			cols = strings.Split(line, "\t")
		case ".csv":
			cols = strings.Split(line, ",")
		}
		if len(cols) != 11 {
			c.LogEventEnd(event_id, "数据格式长度错误", -1, time.Now().Unix())
			return
		}
		cost, err := strconv.Atoi(cols[10])
		if err != nil {
			c.Logger.Println(err)
			c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
			return
		}
		cols[10] = fmt.Sprintf("%v", float64(cost)/1000000)
		args = append(args, InterfaceSlice(cols))

	}
	table_name := "report_base"
	columns := []string{"vendor_id", "user_id", "campaign_id",
		"ad_id", "report_date", "hour", "win", "imp", "clk", "eimp", "cost"}
	//if err := c.CopyReportToDb(table_name, file_ext, root_path+filename, columns); err != nil {
	//	c.Logger.Println(err)
	//	c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
	//	return
	//}
	if err := corm.PgCopy(table_name, args, columns); err != nil {
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}

	c.LogEventEnd(event_id, "", 1, time.Now().Unix())
	c.JsonBase(w, nil)
}
func (c *Handler) ReportGeoImport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Url *string `form:"url"`
	}
	if err := c.BindQuery(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Url == nil {
		c.JsonError(w, my_context.ErrBadValid, nil)
		return
	}
	event_id := uuid.NewV4().String()
	// 事物开始
	c.LogEventStart("基本报表导入", event_id, *req.Url, time.Now().Unix())

	resp, err := http.Get(*req.Url)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	defer resp.Body.Close()
	// 存储文件
	file_ext := filepath.Ext(*req.Url)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + path.Base(*req.Url)
	file, err := os.OpenFile(root_path+filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	lines, err := myfile.ReadLines(root_path + filename)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}

	args := make([][]interface{}, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		cols := make([]string, 0)
		switch file_ext {
		case ".txt":
			cols = strings.Split(line, "\t")
		case ".csv":
			cols = strings.Split(line, ",")
		}
		if len(cols) != 13 {
			c.LogEventEnd(event_id, "数据格式长度错误", -1, time.Now().Unix())
			return
		}
		cost, err := strconv.Atoi(cols[12])
		if err != nil {
			c.Logger.Println(err)
			c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
			return
		}
		cols[12] = fmt.Sprintf("%v", float64(cost)/1000000)
		args = append(args, InterfaceSlice(cols))

	}
	table_name := "report_geo"
	columns := []string{"vendor_id", "user_id", "campaign_id",
		"ad_id", "country_code", "province_code", "city_code", "report_date", "win", "imp", "clk", "eimp", "cost"}
	//if err := c.CopyReportToDb(table_name, file_ext, root_path+filename, columns); err != nil {
	//	c.Logger.Println(err)
	//	c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
	//	return
	//}
	if err := corm.PgCopy(table_name, args, columns); err != nil {
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}

	c.LogEventEnd(event_id, "", 1, time.Now().Unix())
	c.JsonBase(w, nil)

}

func (c *Handler) CopyReportToDb(table_name, file_ext, file_path string, columns []string) error {
	var exec_str string
	switch file_ext {
	case ".csv":
		exec_str = fmt.Sprintf("COPY %s(%s) from '%s' DELIMITER ',' ",
			table_name, strings.Join(columns, ","), file_path)
	case ".txt":
		exec_str = fmt.Sprintf("COPY %s(%s) from '%s' DELIMITER E'\t' ",
			table_name, strings.Join(columns, ","), file_path)
	default:
		return errors.New("文件格式不支持")
	}
	if _, err := c.Pgx.Exec(exec_str); err != nil {
		return err
	}

	return nil
}
func (c *Handler) ReportDevProfitImport(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Url *string `form:"url"`
	}
	if err := c.BindQuery(r, &req); err != nil {
		c.JsonError(w, my_context.ErrBadValid, err)
		return
	}
	if req.Url == nil {
		c.JsonError(w, my_context.ErrBadValid, nil)
		return
	}
	event_id := uuid.NewV4().String()
	// 事物开始
	c.LogEventStart("开发者收益报表导入", event_id, *req.Url, time.Now().Unix())
	resp, err := http.Get(*req.Url)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	defer resp.Body.Close()
	// 存储文件
	file_ext := filepath.Ext(*req.Url)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + path.Base(*req.Url)
	file, err := os.OpenFile(root_path+filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	lines, err := myfile.ReadLines(root_path + filename)
	if err != nil {
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		c.JsonError(w, my_context.ErrDefault, err)
		return
	}
	table_name := "report_dev_profit"
	columns := []string{
		"app_key", "report_date", "bidding", "share",
	}
	//if err := c.CopyReportToDb(table_name, file_ext, root_path+filename, columns); err != nil {
	//	c.Logger.Println(err)
	//	c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
	//	return
	//}
	tx, err := c.Pgx.Begin()
	if err != nil {
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}
	args := make([][]interface{}, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		cols := make([]string, 0)
		switch file_ext {
		case ".txt":
			cols = strings.Split(line, "\t")
		case ".csv":
			cols = strings.Split(line, ",")
		}
		if len(cols) != 4 {
			tx.Rollback()
			c.LogEventEnd(event_id, "数据格式错误", -1, time.Now().Unix())
			return
		}
		bidding, err := strconv.Atoi(cols[2])
		if err != nil {
			c.Logger.Println(err)
			c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
			return
		}
		cols[2] = fmt.Sprintf("%v", float64(bidding)/1000000)
		share, err := strconv.Atoi(cols[3])
		if err != nil {
			c.Logger.Println(err)
			c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
			return
		}
		cols[3] = fmt.Sprintf("%v", float64(share)/1000000)
		args = append(args, InterfaceSlice(cols))
		_, err2 := tx.Exec(`update account_balance set balance=balance+$1  where app_key=$2`, float64(bidding+share)/1000000, cols[0])
		if err2 != nil {
			tx.Rollback()
			c.Logger.Println(err2)
			c.LogEventEnd(event_id, fmt.Sprintf("%v", err2), -1, time.Now().Unix())
			return
		}

	}
	if err := corm.PgCopy(table_name, args, columns); err != nil {
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.Logger.Println(err)
		c.LogEventEnd(event_id, fmt.Sprintf("%v", err), -1, time.Now().Unix())
		return
	}
	c.LogEventEnd(event_id, "", 1, time.Now().Unix())
	c.JsonBase(w, nil)
}

func InterfaceSlice(slice interface{}) []interface{} {
	switch slice.(type) {
	case []interface{}:
		return slice.([]interface{})
	default:
		s := reflect.ValueOf(slice)
		if s.Kind() != reflect.Slice {
			panic("InterfaceSlice() given a non-slice type")
		}
		ret := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}
		return ret
	}

}
