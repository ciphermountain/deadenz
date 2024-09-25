package components

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ciphermountain/deadenz/pkg/opts"
)

type CharacterSpawnEvent struct {
	character Character
	lang      LanguagePack
}

func NewCharacterSpawnEvent(character Character, opts ...opts.Option) *CharacterSpawnEvent {
	lang := &language{}

	for _, opt := range opts {
		opt(lang)
	}

	return &CharacterSpawnEvent{
		character: character,
		lang:      lang.lang}
}

func (e CharacterSpawnEvent) String() string {
	return strings.ReplaceAll(e.lang.SpawninPattern, "{{character}}", e.character.Name)
}

func (e CharacterSpawnEvent) Type() CharacterType {
	return e.character.Type
}

func (e CharacterSpawnEvent) Name() string {
	return e.character.Name
}

func (e CharacterSpawnEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonCharacterSpawnEvent{
		Type:      "spawnin-event",
		Character: e.character,
	})
}

func (e *CharacterSpawnEvent) UnmarshalJSON(bts []byte) error {
	var base jsonCharacterSpawnEvent
	if err := json.Unmarshal(bts, &base); err != nil {
		return err
	}

	if base.Type != "spawnin-event" {
		return fmt.Errorf("%w: %s; expected spawnin-event", ErrInvalidEventType, base.Type)
	}

	*e = CharacterSpawnEvent{character: base.Character}

	return nil
}

type jsonCharacterSpawnEvent struct {
	Type      string    `json:"type"`
	Character Character `json:"character"`
}

type EarnedXPEvent struct {
	xp   uint
	lang LanguagePack
}

func NewEarnedXPEvent(xp uint, opts ...opts.Option) *EarnedXPEvent {
	lang := &language{}

	for _, opt := range opts {
		opt(lang)
	}

	return &EarnedXPEvent{xp: xp, lang: lang.lang}
}

func (e EarnedXPEvent) String() string {
	return strings.ReplaceAll(e.lang.EarnedXPPattern, "{{amount}}", strconv.FormatUint(uint64(e.xp), 10))
}

type EarnedTokenEvent struct {
	xp   uint
	lang LanguagePack
}

func NewEarnedTokenEvent(xp uint, opts ...opts.Option) *EarnedTokenEvent {
	lang := &language{}

	for _, opt := range opts {
		opt(lang)
	}

	return &EarnedTokenEvent{xp: xp, lang: lang.lang}
}

func (e EarnedTokenEvent) String() string {
	return strings.ReplaceAll(e.lang.EarnedTokenPattern, "{{amount}}", strconv.FormatUint(uint64(e.xp), 10))
}
