// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

import React from 'react'

import Dropdown from '../dropdown'

import getProfileDropdownDriver from './index_driver'

const userName = 'kschiffer'
const nullAction = function() {
  return null
}
const children = (
  <React.Fragment>
    <Dropdown.Item title="Profile Settings" icon="settings" link="/profile-settings" />
    <Dropdown.Item title="Logout" icon="power_settings_new" action={nullAction} />
  </React.Fragment>
)

describe('<ProfileDropdown />', function() {
  let driver = null

  beforeEach(function() {
    driver = getProfileDropdownDriver()
  })

  describe('when in initial state', function() {
    beforeEach(function() {
      driver.when.created({ userName, children })
    })
    it('matches snapshot', function() {
      expect(driver.component).toMatchSnapshot()
    })
    it('does not have the dropdown open by default', function() {
      expect(driver.is.dropdownOpen()).toBeFalsy()
    })
    it('opens dropdown on click', function() {
      driver.when.toggledDropdown()
      expect(driver.is.dropdownOpen()).toBeTruthy()
    })

    describe('when it has dropdown open', function() {
      beforeEach(function() {
        driver.when.toggledDropdown()
      })
      it('matches snapshot', function() {
        expect(driver.component).toMatchSnapshot()
      })
      it('closes dropdown on click', function() {
        driver.when.toggledDropdown()
        expect(driver.is.dropdownOpen()).toBeFalsy()
      })
    })
  })
})
