// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crypto

import (
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/pkg/types"
	"go.thethings.network/lorawan-stack/pkg/util/test/assertions/should"
)

func TestKeyDerivation(t *testing.T) {
	a := assertions.New(t)

	a.So(reverse([]byte{1, 2, 3, 4}), should.Resemble, []byte{4, 3, 2, 1})

	key := types.AES128Key{0xBE, 0xC4, 0x99, 0xC6, 0x9E, 0x9C, 0x93, 0x9E, 0x41, 0x3B, 0x66, 0x39, 0x61, 0x63, 0x6C, 0x61}
	dn := types.DevNonce{0x73, 0x69}
	nid := types.NetID{0x02, 0x01, 0x01}
	jn := types.JoinNonce{0xAE, 0x3B, 0x1C}
	joinEUI := types.EUI64{0x00, 0x00, 0x00, 0x12, 0x23, 0x22, 0x42, 0x42}
	devEUI := types.EUI64{0x00, 0x42, 0x00, 0x42, 0x13, 0x42, 0x43, 0x41}

	appSKey := DeriveAppSKey(key, jn, joinEUI, dn)
	a.So(appSKey, should.Equal, types.AES128Key{0xB0, 0x35, 0xCC, 0xE2, 0x4E, 0x4A, 0x04, 0x0E, 0xB5, 0xCE, 0xCB, 0xF2, 0x17, 0x24, 0x41, 0x2D})

	fNwkSIntKey := DeriveFNwkSIntKey(key, jn, joinEUI, dn)
	a.So(fNwkSIntKey, should.Equal, types.AES128Key{0x37, 0x90, 0x84, 0xE7, 0xCE, 0x22, 0xFF, 0x19, 0x1B, 0xFF, 0x4B, 0x77, 0x53, 0x6F, 0x2A, 0xA3})

	sNwkSIntKey := DeriveSNwkSIntKey(key, jn, joinEUI, dn)
	a.So(sNwkSIntKey, should.Equal, types.AES128Key{0x63, 0xC5, 0x93, 0x09, 0xD5, 0x34, 0x85, 0xBC, 0x51, 0x64, 0xDB, 0xF7, 0x16, 0x27, 0xAE, 0xB9})

	nwkSEncKey := DeriveNwkSEncKey(key, jn, joinEUI, dn)
	a.So(nwkSEncKey, should.Equal, types.AES128Key{0xCE, 0x07, 0xA0, 0x09, 0xA3, 0x97, 0x0A, 0xC0, 0x51, 0x9A, 0x09, 0x9E, 0xD5, 0x3E, 0x55, 0x0B})

	appSKey = DeriveLegacyAppSKey(key, jn, nid, dn)
	a.So(appSKey, should.Equal, types.AES128Key{0x8C, 0x1E, 0x05, 0x43, 0xA2, 0x29, 0x08, 0x8D, 0xE6, 0xF8, 0x4E, 0x74, 0xBB, 0x46, 0xBD, 0x62})

	nwkSKey := DeriveLegacyNwkSKey(key, jn, nid, dn)
	a.So(nwkSKey, should.Equal, types.AES128Key{0x0D, 0xB9, 0x24, 0xEE, 0x6A, 0xF9, 0x06, 0x98, 0xE0, 0x5F, 0xC7, 0xCE, 0x48, 0x30, 0x3C, 0x01})

	jsIntKey := DeriveJSIntKey(key, devEUI)
	a.So(jsIntKey, should.Equal, types.AES128Key{0xCB, 0x8E, 0x9E, 0xE1, 0x14, 0xC3, 0xFB, 0x5C, 0xE4, 0x72, 0xA5, 0xAA, 0x13, 0xAB, 0x18, 0x55})

	jsEncKey := DeriveJSEncKey(key, devEUI)
	a.So(jsEncKey, should.Equal, types.AES128Key{0xBB, 0x71, 0x1E, 0xEF, 0xB9, 0x82, 0x9B, 0x4A, 0x75, 0x86, 0x6F, 0x86, 0x16, 0xBA, 0xCD, 0x6D})
}
