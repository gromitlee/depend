package examples

import (
	"testing"

	"github.com/gromitlee/depend"
	"github.com/gromitlee/depend/rel"
	"gorm.io/gorm"
)

const (
	resProject  rel.ResType = 1
	resTag      rel.ResType = 2
	resDataset  rel.ResType = 3
	resTrain    rel.ResType = 4
	resTest     rel.ResType = 5
	resSdk      rel.ResType = 6
	resPipeline rel.ResType = 7
	resInfer    rel.ResType = 8
)

func TestDepend(t *testing.T) {
	db := getDB(dbMysql, dbName)
	if err := initDepend(db); err != nil {
		t.Fatal(err)
	}

	res201 := rel.Res{ID: "tag_201", Typ: resTag}
	res202 := rel.Res{ID: "tag_202", Typ: resTag}
	res203 := rel.Res{ID: "tag_203", Typ: resTag}
	res400 := rel.Res{ID: "train_400", Typ: resTrain}
	res700 := rel.Res{ID: "pipeline_700", Typ: resPipeline}

	// 301
	res301 := rel.Res{ID: "dataset_301", Typ: resDataset}
	if err := depend.AddRelation(db, res301, res400); err != nil {
		t.Fatal(err)
	}
	t.Log(depend.CheckOp(db, res301, rel.OpDelete))
	t.Log(depend.CheckOp(db, res301, rel.OpEdit))
	if err := depend.CheckRelation(db, res301, res400); err != nil {
		t.Fatal(err)
	}

	// 302
	res302 := rel.Res{ID: "dataset_302", Typ: resDataset}
	if err := depend.AddRelation(db, res302, res700); err != nil {
		t.Fatal(err)
	}
	t.Log(depend.CheckOp(db, res302, rel.OpDelete))
	if err := depend.CheckOp(db, res302, rel.OpEdit); err != nil {
		t.Log(err)
	}
	if relations, err := depend.GetRelations(db, res302); err != nil {
		t.Fatal(err)
	} else {
		t.Log(relations[0])
	}

	// 303
	res303 := rel.Res{ID: "dataset_303", Typ: resDataset}
	if err := depend.AddRelations(db, res303, []rel.Res{res400, res700}); err != nil {
		t.Fatal(err)
	}
	if err := depend.DelRelation(db, res303, res400); err != nil {
		t.Fatal()
	}
	t.Log(depend.CheckOp(db, res303, rel.OpDelete))
	if err := depend.CheckOp(db, res303, rel.OpEdit); err != nil {
		t.Log(err)
	}
	if depends, err := depend.GetDependents(db, res700); err != nil {
		t.Log(err)
	} else {
		t.Log(depends[0], depends[1])
	}

	// 304
	res304 := rel.Res{ID: "dataset_304", Typ: resDataset}
	if err := depend.AddDepends(db, []rel.Res{res201, res202, res203}, res304); err != nil {
		t.Fatal(err)
	}
	if depends, err := depend.GetDependents(db, res304); err != nil {
		t.Log(err)
	} else {
		t.Log(depends[0], depends[1], depends[2])
	}

	// clean
	if err := depend.DelRes(db, res700); err != nil {
		t.Fatal(err)
	}
	if err := depend.DelRes(db, res400); err != nil {
		t.Fatal(err)
	}
	if err := depend.DelRes(db, res304); err != nil {
		t.Fatal(err)
	}
}

func initDepend(db *gorm.DB) error {
	if err := depend.Init(db); err != nil {
		return err
	}

	// project
	if err := depend.Register(resProject, rel.OpDelete, nil); err != nil {
		return err
	}

	// tag
	if err := depend.Register(resTag, rel.OpDelete, nil); err != nil {
		return err
	}

	// dataset
	if err := depend.Register(resDataset, rel.OpDelete, nil); err != nil {
		return err
	}
	if err := depend.Register(resDataset, rel.OpEdit, []rel.ResType{resPipeline}); err != nil {
		return err
	}

	// train
	if err := depend.Register(resTrain, rel.OpDelete, nil); err != nil {
		return err
	}
	if err := depend.Register(resTrain, rel.OpEdit, []rel.ResType{resTest, resSdk, resPipeline}); err != nil {
		return err
	}

	// test
	if err := depend.Register(resTest, rel.OpDelete, nil); err != nil {
		return err
	}
	if err := depend.Register(resTest, rel.OpEdit, []rel.ResType{resSdk, resPipeline}); err != nil {
		return err
	}

	// sdk
	if err := depend.Register(resSdk, rel.OpDelete, nil); err != nil {
		return err
	}
	if err := depend.Register(resTest, rel.OpEdit, []rel.ResType{resInfer}); err != nil {
		return err
	}

	// pipeline
	if err := depend.Register(resPipeline, rel.OpDelete, nil); err != nil {
		return err
	}

	// infer
	if err := depend.Register(resInfer, rel.OpDelete, nil); err != nil {
		return err
	}

	return nil
}
