package main

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/ktr0731/go-fuzzyfinder"
)

func main() {

	// あふのフォルダ履歴取得
	dirs, err := histories()
	failOnError(err)

	// 検索
	idx, _ := fuzzyfinder.Find(dirs, func(i int) string {
		return dirs[i]
	})

	// フォルダ変更
	failOnError(excd(dirs[idx]))
}

// あふのフォルダ履歴を取得
func histories() ([]string, error) {
	_ = ole.CoInitialize(0)
	unknown, err := oleutil.CreateObject("afxw.obj")
	if err != nil {
		return nil, err
	}
	defer unknown.Release()
	afxw, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}

	var dirs []string

	// あふの左右のウィンドウ分ループ
	for win := 0; win < 2; win++ {
		count := oleutil.MustCallMethod(afxw, "HisDirCount", win).Value().(int32)
		for i := 0; i < int(count); i++ {
			dir := oleutil.MustCallMethod(afxw, "HisDir", win, i)
			dirs = append(dirs, fmt.Sprint(dir.Value()))
		}
	}

	return dirs, nil
}

// あふのフォルダ変更
func excd(path string) error {
	_ = ole.CoInitialize(0)
	unknown, err := oleutil.CreateObject("afxw.obj")
	if err != nil {
		return err
	}
	defer unknown.Release()

	afxw, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}

	oleutil.MustCallMethod(afxw, "Exec", fmt.Sprintf("&EXCD -P%s", path))
	return nil
}

// エラー処理
func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
