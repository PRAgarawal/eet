package grouping

import (
	"strconv"
	"testing"

	"github.com/PRAgarawal/eet/eet"
)

func TestSmallestGroup(t *testing.T) {
	groupSize := 5
	mGrouping := MemberGrouping{}
	mGrouping.SetMembers(makeGroup(groupSize))

	groups := mGrouping.RandomlyGroup()

	if len(groups) != 1 {
		t.Fatalf("Expected 1 group, got %d groups", len(groups))
	}

	if len(groups[0]) != groupSize {
		t.Errorf("Expected group of size %d, got size %d", groupSize, len(groups[0]))
	}

	checkForRandomness(t, groups)
}

func TestSmallGroup(t *testing.T) {
	groupSize := 10
	splitGroupSize := 5
	mGrouping := MemberGrouping{}
	mGrouping.SetMembers(makeGroup(groupSize))

	groups := mGrouping.RandomlyGroup()

	if len(groups) != 2 {
		t.Fatalf("Expected 2 groups, got %d groups", len(groups))
	}

	if len(groups[0]) != splitGroupSize {
		t.Errorf("Expected group 0 size %d, got size %d", splitGroupSize, len(groups[0]))
	}

	if len(groups[1]) != splitGroupSize {
		t.Errorf("Expected group 1 size %d, got size %d", splitGroupSize, len(groups[1]))
	}

	checkForRandomness(t, groups)
}

// Remainder after "ideal" group size is less than "minimum" group size
func TestGroupRemainder1(t *testing.T) {
	groupSize := 27
	firstGroupsSize := idealGroupSize
	nextSize := 5
	finalSize := 4
	mGrouping := MemberGrouping{}
	mGrouping.SetMembers(makeGroup(groupSize))

	groups := mGrouping.RandomlyGroup()

	if len(groups) != 5 {
		t.Fatalf("Expected 5 groups, got %d groups", len(groups))
	}

	for i := 0; i < 3; i++ {
		if len(groups[i]) != firstGroupsSize {
			t.Errorf("Expected group %d to be of size %d, got size %d", i, idealGroupSize, len(groups[i]))
		}
	}

	if len(groups[3]) != nextSize {
		t.Errorf("Expected group 3 size %d, got size %d", nextSize, len(groups[5]))
	}

	if len(groups[4]) != finalSize {
		t.Errorf("Expected group 4 size %d, got size %d", finalSize, len(groups[4]))
	}

	checkForRandomness(t, groups)
}

// Remainder after "ideal" group size is greater than or equal to "minimum" group size
func TestGroupRemainder2(t *testing.T) {
	groupSize := 23
	firstGroupsSize := idealGroupSize
	finalSize := 5
	mGrouping := MemberGrouping{}
	mGrouping.SetMembers(makeGroup(groupSize))

	groups := mGrouping.RandomlyGroup()

	if len(groups) != 4 {
		t.Fatalf("Expected 4 groups, got %d groups", len(groups))
	}

	for i := 0; i < 3; i++ {
		if len(groups[i]) != firstGroupsSize {
			t.Errorf("Expected group %d to be of size %d, got size %d", i, idealGroupSize, len(groups[i]))
		}
	}

	if len(groups[3]) != finalSize {
		t.Errorf("Expected group 4 size %d, got size %d", finalSize, len(groups[3]))
	}

	checkForRandomness(t, groups)
}

// Exact multiple of "ideal" group size
func TestGroupRemainder3(t *testing.T) {
	groupSize := 30
	firstGroupsSize := idealGroupSize
	mGrouping := MemberGrouping{}
	mGrouping.SetMembers(makeGroup(groupSize))

	groups := mGrouping.RandomlyGroup()

	if len(groups) != 5 {
		t.Fatalf("Expected 5 groups, got %d groups", len(groups))
	}

	for i := 0; i < 5; i++ {
		if len(groups[i]) != firstGroupsSize {
			t.Errorf("Expected group %d to be of size %d, got size %d", i, idealGroupSize, len(groups[i]))
		}
	}

	checkForRandomness(t, groups)
}

// Kinda shady randomness test
func checkForRandomness(t *testing.T, groups [][]*eet.MeetingMember) {
	groupIds := make(map[string]string)
	prev := -1
	successionCount := 0
	totalCount := 0

	for _, group := range groups {
		for _, member := range group {
			totalCount++

			if groupIds[member.SlackUserId] != "" {
				t.Error("Expected unique slack user ID, but found a repeat")
			}

			groupIds[member.SlackUserId] = member.SlackUserId

			id, _ := strconv.Atoi(member.SlackUserId)
			if id == (prev + 1) {
				successionCount++
			}
			prev = id
		}
	}

	if successionCount > totalCount/4 {
		// Not sufficiently random... I guess?
		t.Error("Expected sufficient randomness in grouping, but got many ordered results")
	}
}

func makeGroup(size int) []*eet.MeetingMember {
	group := make([]*eet.MeetingMember, size)

	for i := range group {
		group[i] = &eet.MeetingMember{
			SlackUserId: strconv.Itoa(i + 1),
		}
	}

	return group
}
