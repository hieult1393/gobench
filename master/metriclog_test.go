package master

import (
	"context"
	"testing"
	"time"

	entApplication "github.com/gobench-io/gobench/ent/application"
	entGraph "github.com/gobench-io/gobench/ent/graph"
	entGroup "github.com/gobench-io/gobench/ent/group"
	entMetric "github.com/gobench-io/gobench/ent/metric"
	"github.com/gobench-io/gobench/metrics"
	"github.com/stretchr/testify/assert"
)

func TestFindCreateGroup(t *testing.T) {
	var err error
	ctx := context.TODO()

	s := seedServer(t)

	s.master.job.app, err = s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	groupRes := new(metrics.FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroup(
		&metrics.FCGroupReq{Name: groupName, AppID: s.master.job.app.ID},
		groupRes))

	// read from db, check with groupRes
	groups, err := s.master.db.Group.Query().Where(
		entGroup.Name(groupName),
		entGroup.HasApplicationWith(
			entApplication.NameEQ("name"),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, groups, 1)
	g := groups[0]
	assert.Equal(t, g.ID, groupRes.ID)

	// call the same RPC, the result should be like before
	groupRes2 := new(metrics.FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroup(
		&metrics.FCGroupReq{Name: groupName, AppID: s.master.job.app.ID},
		groupRes2))
	assert.Equal(t, groupRes, groupRes2)
}

func TestFindCreateGraph(t *testing.T) {
	var err error
	ctx := context.TODO()

	s := seedServer(t)

	s.master.job.app, err = s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	groupRes := new(metrics.FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroup(
		&metrics.FCGroupReq{Name: groupName, AppID: s.master.job.app.ID},
		groupRes))

	// create new graph
	graphReq := &metrics.FCGraphReq{
		Title:   "HTTP Response",
		Unit:    "N",
		GroupID: groupRes.ID,
	}
	graphRes := new(metrics.FCGraphRes)
	assert.Nil(t, s.master.FindCreateGraph(graphReq, graphRes))

	// read from db, check with groupRes
	graphs, err := s.master.db.Graph.Query().Where(
		entGraph.TitleEQ(graphReq.Title),
		entGraph.HasGroupWith(
			entGroup.IDEQ(groupRes.ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, graphs, 1)
	g := graphs[0]
	assert.Equal(t, g.ID, graphRes.ID)

	// call the same RPC, the result should be like before
	graphRes2 := new(metrics.FCGraphRes)
	assert.Nil(t, s.master.FindCreateGraph(graphReq, graphRes2))
	assert.Equal(t, graphRes, graphRes2)
}

func TestFindCreateMetric(t *testing.T) {
	var err error
	ctx := context.TODO()

	s := seedServer(t)

	s.master.job.app, err = s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	// create new group
	groupRes := new(metrics.FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroup(
		&metrics.FCGroupReq{Name: groupName, AppID: s.master.job.app.ID},
		groupRes))

	// create new graph
	graphReq := &metrics.FCGraphReq{
		Title:   "HTTP Response",
		Unit:    "N",
		GroupID: groupRes.ID,
	}
	graphRes := new(metrics.FCGraphRes)
	assert.Nil(t, s.master.FindCreateGraph(graphReq, graphRes))

	// create new metric
	metricReq := &metrics.FCMetricReq{
		Title:   ".http_ok",
		Type:    metrics.Counter,
		GraphID: graphRes.ID,
	}
	metricRes := new(metrics.FCMetricRes)
	assert.Nil(t, s.master.FindCreateMetric(metricReq, metricRes))

	// call the same RPC, the result should be like before
	metricRes2 := new(metrics.FCGraphRes)
	assert.Nil(t, s.master.FindCreateGraph(graphReq, metricRes2))
	assert.Equal(t, graphRes, metricRes2)

	// read from db, check with groupRes
	metrics, err := s.master.db.Metric.Query().Where(
		entMetric.TitleEQ(metricReq.Title),
		entMetric.HasGraphWith(
			entGraph.IDEQ(graphRes.ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, metrics, 1)
	m := metrics[0]
	assert.Equal(t, m.ID, metricRes.ID)
}
