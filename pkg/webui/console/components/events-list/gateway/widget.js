// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

import Events from '@ttn-lw/components/events-list'

import PropTypes from '@ttn-lw/lib/prop-types'

import renderGatewayEvent from './render-event'

const renderEvent = event => renderGatewayEvent(event, true)

const GatewayEvents = props => {
  const { events, gtwId, toAllUrl } = props

  return (
    <Events.Widget events={events} renderEvent={renderEvent} toAllUrl={toAllUrl} entityId={gtwId} />
  )
}

GatewayEvents.propTypes = {
  events: PropTypes.arrayOf(PropTypes.event).isRequired,
  gtwId: PropTypes.string.isRequired,
  toAllUrl: PropTypes.string.isRequired,
}

export default GatewayEvents
