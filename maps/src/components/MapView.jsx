import React from 'react';
import { MapContainer, TileLayer, Marker, Popup, useMap } from 'react-leaflet';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';



// this file handles the Map Display
//It renders the following : 
// The main map container using OpenStreetMap tiles

// Markers for all locations stored in state

// Popups with location details when markers are clicked

// Handles map click events (passed from parent)

// Fix for default marker icons

// Uses react-leaflet (React wrapper for Leaflet.js)

// Displays OpenStreetMap tiles

// Shows markers for all saved locations

// Supports popups with location details


delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon-2x.png',
  iconUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png',
  shadowUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-shadow.png',
});

// The Three Images Being Loaded
// iconUrl: Standard green marker (25×41 pixels)

// https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png

// iconRetinaUrl: High-DPI version (50×82 pixels)

// https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon-2x.png

// shadowUrl: Drop shadow for the marker (41×41 pixels)

// https://unpkg.com/leaflet@1.7.1/dist/images/marker-shadow.png
//Leaflet is a lightweight, open-source JavaScript library for creating interactive maps in web applications
// Leaflet Creates interactive maps with panning/zooming

// Displays map tiles from providers (OpenStreetMap, Mapbox, etc.)

// Adds markers, shapes, and popups to maps

// Handles user interactions (clicks, drags, etc.)

const MapView = ({ center, zoom, markers, onMapClick }) => {
  return (
    <MapContainer 
      center={center} 
      zoom={zoom} 
      style={{ height: '50vh', width: '1000px' }}
      onClick={onMapClick}
    >
      {/* OpenFreeMap Tile Layer */}
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      />
      {markers.map((marker, index) => (
        <Marker key={index} position={marker.position}>
          <Popup>
            <div>
              <h3>{marker.title}</h3>
              <p>{marker.description}</p>
              {marker.image && <img src={marker.image} alt={marker.title} width="100" />}
            </div>
          </Popup>
        </Marker>
      ))}
    </MapContainer>
  );
};

export default MapView;