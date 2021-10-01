// Copyright (c) 2020 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package schema

import "github.com/contaim/spec/parser"

// VariablesDefinition allows registration of the vars definition block using
// NewWithSubset.
func VariablesDefinition() parser.NamedBlockDefinition {
	return &variablesDef{}
}

// HideBannerDefinition allows registration of the hide_banner definition block
// using NewWithSubset.
func HideBannerDefinition() parser.NamedBlockDefinition {
	return &hideBannerDef{}
}

// HideServingDefinition allows registration of the hide_banner definition block
// using NewWithSubset.
func HideServingDefinition() parser.NamedBlockDefinition {
	return &hideServingDef{}
}

// UIDirDefinition allows registration of the ui_dir definition block
// using NewWithSubset.
func UIDirDefinition() parser.NamedBlockDefinition {
	return &uiDirDef{}
}

// EncryptionKeyDefinition allows registration of the encryption_key definition
// block using NewWithSubset.
func EncryptionKeyDefinition() parser.NamedBlockDefinition {
	return &encryptionKeyDef{}
}

// EventsDefinition allows registration of the events block using
// NewWithSubset.
func EventsDefinition() parser.NamedBlockDefinition {
	return &eventsDef{}
}

// DatabaseDefinition allows registration of the ent block using
// NewWithSubset.
func DatabaseDefinition() parser.NamedBlockDefinition {
	return &databaseDef{}
}

func DeveloperDefinition() parser.NamedBlockDefinition {
	return &developerDef{}
}

func ClusterDefinition() parser.NamedBlockDefinition {
	return &clusterDef{}
}

func HTTPListenerDefinition() parser.NamedBlockDefinition {
	return &httpDef{}
}

func GameListenerDefinition() parser.NamedBlockDefinition {
	return &gameDef{}
}

func AuthDefinition() parser.NamedBlockDefinition {
	return &authDef{}
}

// Schema defines the blocks and their respecify BlockDefinition by creating
// an instance of parser.BlockDefinitions.
func Schema() parser.NamedBlockDefinitions {
	return parser.NamedBlockDefinitions{

		// variables must be processed first
		VariablesDefinition(),

		// encryption_key attribute
		EncryptionKeyDefinition(),

		// ui_dir attribute
		UIDirDefinition(),

		// hide_banner attribute
		HideBannerDefinition(),

		// hide_serving attribute
		HideServingDefinition(),

		// listener blocks
		HTTPListenerDefinition(),
		GameListenerDefinition(),

		// // settings block
		// SettingsDefinition(),

		// auth block
		AuthDefinition(),

		// database block
		DatabaseDefinition(),

		// pubsub block
		EventsDefinition(),

		// developer block
		DeveloperDefinition(),

		// cluster block
		ClusterDefinition(),
	}
}
