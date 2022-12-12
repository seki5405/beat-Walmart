import React from "react";
import {useState} from 'react';
import { geoCentroid } from "d3-geo";
import {
  ComposableMap,
  Geographies,
  Geography,
  Marker,
  Annotation
} from "react-simple-maps";

import allStates from "./data/allstates.json";
import data from "./data.json";

const geoUrl = "https://cdn.jsdelivr.net/npm/us-atlas@3/states-10m.json";
const url = "http://localhost:30001/poll"

const offsets = {
  VT: [50, -8],
  NH: [34, 2],
  MA: [30, -1],
  RI: [28, 2],
  CT: [35, 10],
  NJ: [34, 1],
  DE: [33, 0],
  MD: [47, 10],
  DC: [49, 21]
};

const state_mapping = {
  "AL": "Alabama",
  "AK": "Alaska",
  "AS": "American Samoa",
  "AZ": "Arizona",
  "AR": "Arkansas",
  "CA": "California",
  "CO": "Colorado",
  "CT": "Connecticut",
  "DE": "Delaware",
  "DC": "District Of Columbia",
  "FM": "Federated States Of Micronesia",
  "FL": "Florida",
  "GA": "Georgia",
  "GU": "Guam",
  "HI": "Hawaii",
  "ID": "Idaho",
  "IL": "Illinois",
  "IN": "Indiana",
  "IA": "Iowa",
  "KS": "Kansas",
  "KY": "Kentucky",
  "LA": "Louisiana",
  "ME": "Maine",
  "MH": "Marshall Islands",
  "MD": "Maryland",
  "MA": "Massachusetts",
  "MI": "Michigan",
  "MN": "Minnesota",
  "MS": "Mississippi",
  "MO": "Missouri",
  "MT": "Montana",
  "NE": "Nebraska",
  "NV": "Nevada",
  "NH": "New Hampshire",
  "NJ": "New Jersey",
  "NM": "New Mexico",
  "NY": "New York",
  "NC": "North Carolina",
  "ND": "North Dakota",
  "MP": "Northern Mariana Islands",
  "OH": "Ohio",
  "OK": "Oklahoma",
  "OR": "Oregon",
  "PW": "Palau",
  "PA": "Pennsylvania",
  "PR": "Puerto Rico",
  "RI": "Rhode Island",
  "SC": "South Carolina",
  "SD": "South Dakota",
  "TN": "Tennessee",
  "TX": "Texas",
  "UT": "Utah",
  "VT": "Vermont",
  "VI": "Virgin Islands",
  "VA": "Virginia",
  "WA": "Washington",
  "WV": "West Virginia",
  "WI": "Wisconsin",
  "WY": "Wyoming"
}

const markers = [
    {
      markerOffset: -30,
      name: "Buenos Aires",
      coordinates: [-58.3816, -34.6037]
    },
    { markerOffset: 15, name: "La Paz", coordinates: [-68.1193, -16.4897] }
  ];

const MapChart = ({ setTooltipContent }) => {
  const [data, setData] = React.useState([]);
  const fetchData = async () => {
    try {
      const response = await fetch(url);
      console.log(response);
      const data = await response.json();
      setData(data);
      console.log(data);
      return data;
    } catch (error) {
      console.log(error);
    }
  };
    
  return (
    <div data-tip="">
    <ComposableMap projection="geoAlbersUsa">
      <Geographies geography={geoUrl}>
        {({ geographies }) => (
          <>
            {geographies.map(geo => (
              <Geography
                key={geo.rsmKey}
                stroke="#FFF"
                geography={geo}
                fill="#DDD"
                onMouseEnter={() => {
                  if (data.length == 0) {
                    fetchData();
                  }
                  console.log(geo.properties.name);
                  
                  var state_arr = [];

                  for (var i = 0; i < data.length; i++) {
                    state_arr.push(state_mapping[data[i].state]);
                  }

                  console.log(state_arr);

                  
                  const cur = data.find(s => state_mapping[s.state] === geo.properties.name);
                  console.log(cur)
                  setTooltipContent(`Item running low : ${cur.low_product}`);
                  
                  }}
                onMouseLeave={() => {
                  setTooltipContent("");
                }}
                style={{
                default: {
                    fill: "#D6D6DA",
                    outline: "none",
                },
                hover: {
                    fill: "#E42",
                    outline: "none"
                },
                pressed: {
                    fill: "#E42",
                    outline: "none"
                }
                }}
              />
            ))}
            {geographies.map(geo => {
              const centroid = geoCentroid(geo);
              const cur = allStates.find(s => s.val === geo.id);
              return (
                <g key={geo.rsmKey + "-name"}>
                  {cur &&
                    centroid[0] > -160 &&
                    centroid[0] < -67 &&
                    (Object.keys(offsets).indexOf(cur.id) === -1 ? (
                      <Marker coordinates={centroid}>
                        <g
                            fill="none"
                            stroke="#FF5533"
                            strokeWidth="2"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            transform="translate(-12, -24)"
                        >
                            <circle cx="12" cy="10" r="3" />
                            <path d="M12 21.7C17.3 17 20 13 20 10a8 8 0 1 0-16 0c0 3 2.7 6.9 8 11.7z" />
                        </g>
                        <text y="2" fontSize={14} textAnchor="middle">
                          {cur.id}
                        </text>
                      </Marker>
                    ) : (
                      <Annotation
                        subject={centroid}
                        dx={offsets[cur.id][0]}
                        dy={offsets[cur.id][1]}
                      >
                        <text x={4} fontSize={14} alignmentBaseline="middle">
                          {cur.id}
                        </text>
                      </Annotation>
                    ))}
                </g>
              );
            })}
          </>
        )}
      </Geographies>
    </ComposableMap>
    </div>
  );
};

export default MapChart;
