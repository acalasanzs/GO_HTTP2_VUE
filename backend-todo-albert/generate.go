package main

import (
	"math/rand"
	"time"
)

// NewNameGenerator creates a new name generator
func NewNameGenerator() *NameGenerator {
	adjectives := []string{
		"Cosmic", "Quantum", "Epic", "Mystic", "Stellar", "Sonic", "Astral", "Crystal",
		"Radiant", "Phantom", "Emerald", "Obsidian", "Arctic", "Azure", "Golden", "Scarlet",
		"Midnight", "Crimson", "Sapphire", "Jade", "Ruby", "Onyx", "Neon", "Laser",
		"Thunder", "Lightning", "Frost", "Flame", "Shadow", "Lunar", "Solar", "Neo",
		"Cyber", "Digital", "Electric", "Hyper", "Ultra", "Mega", "Super", "Primal",
		"Ancient", "Eternal", "Infinite", "Immortal", "Legendary", "Mythic", "Heroic", "Epic",
		"Savage", "Fierce", "Majestic", "Royal", "Noble", "Imperial", "Rogue", "Wild",
	}

	nouns := []string{
		"Phoenix", "Dragon", "Titan", "Nexus", "Spectre", "Phantom", "Oracle", "Sentinel",
		"Voyager", "Guardian", "Paladin", "Raven", "Wolf", "Falcon", "Eagle", "Shark",
		"Panther", "Tiger", "Lion", "Cobra", "Viper", "Vector", "Matrix", "Fusion",
		"Nebula", "Galaxy", "Comet", "Nova", "Pulsar", "Quasar", "Zenith", "Horizon",
		"Vortex", "Tempest", "Cyclone", "Inferno", "Tundra", "Oasis", "Cascade", "Summit",
		"Forge", "Bastion", "Citadel", "Fortress", "Sanctum", "Vault", "Monolith", "Spire",
		"Blade", "Shield", "Hammer", "Scepter", "Crown", "Tome", "Relic", "Artifact",
		"Construct", "Engine", "Core", "Beacon", "Signal", "Echo", "Whisper", "Enigma",
	}

	source := rand.NewSource(time.Now().UnixNano())
	return &NameGenerator{
		adjectives: adjectives,
		nouns:      nouns,
		rng:        rand.New(source),
	}
}
