package grouping

import (
	"math/rand"

	"github.com/PRAgarawal/eet/eet"
)

const (
	minGroupSize = 5
	idealGroupSize = 6
	maxGroupSize = 8
)

type MemberGrouping struct {
	members []*eet.MeetingMember
	idealGroupCount int
	remainder int
	remainderGroupCount int
	groups [][]*eet.MeetingMember
}

func (g *MemberGrouping) SetMembers(members []*eet.MeetingMember) {
	g.members = members
}

func (g *MemberGrouping) GetGroups() [][]*eet.MeetingMember {
	return g.groups
}

func (g *MemberGrouping) RandomlyGroup() [][]*eet.MeetingMember {
	// shuffle members
	for i := range g.members {
		j := rand.Intn(i + 1)
		g.members[i], g.members[j] = g.members[j], g.members[i]
	}

	// Small batch
	if len(g.members) <= maxGroupSize {
		return [][]*eet.MeetingMember{g.members}
	}

	g.setGroupCounts()

	g.groups = make([][]*eet.MeetingMember, g.idealGroupCount + g.remainderGroupCount)

	g.divideIdeallySizedGroups()
	g.divideRemainderGroups()

	return g.groups
}

func (g *MemberGrouping) setGroupCounts() {
	g.idealGroupCount = len(g.members)/idealGroupSize
	g.remainder = len(g.members)%idealGroupSize

	if g.remainder < minGroupSize {
		g.idealGroupCount -= 1
		g.remainder += idealGroupSize
	}

	if g.remainder > maxGroupSize {
		g.remainderGroupCount = 2
	} else if g.remainder > 0 {
		g.remainderGroupCount = 1
	}
}

func (g *MemberGrouping) divideIdeallySizedGroups() {
	for i := 0; i < g.idealGroupCount; i++ {
		g.groups[i] = make([]*eet.MeetingMember, idealGroupSize)
		for j := 0; j < idealGroupSize; j++ {
			g.groups[i][j] = g.members[i * idealGroupSize + j]
		}
	}
}

func (g *MemberGrouping) divideRemainderGroups() {
	if g.remainder > maxGroupSize {
		g.groups[g.idealGroupCount] = make([]*eet.MeetingMember, minGroupSize)
		g.groups[g.idealGroupCount + 1] = make([]*eet.MeetingMember, g.remainder - minGroupSize)
	} else {
		g.groups[g.idealGroupCount] = make([]*eet.MeetingMember, g.remainder)
	}

	for i := 0; i < g.remainder; i++ {
		offset := g.idealGroupCount * idealGroupSize + i

		if i < minGroupSize || g.remainder <= maxGroupSize {
			g.groups[g.idealGroupCount][i] = g.members[offset]
		} else {
			g.groups[g.idealGroupCount + 1][i - minGroupSize] = g.members[offset]
		}
	}
}
