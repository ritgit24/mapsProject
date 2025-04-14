import React, { useState } from 'react';
import axios from 'axios';
import './search.css'

// it renders the search Interface:

// Text input for place search

// Loading indicator during API calls

// Dropdown with search results

// Map Container (L.map()) - The base layer

// Tile Layers - Background map imagery

// Markers - Location indicators

// Vector Layers - Lines, polygons, circles

// Controls - Zoom buttons, layer switchers

const SearchAndAdd = ({ onSearch}) => {
  const [searchResults, setSearchResults] = useState([]);
//   const [selectedLocation, setSelectedLocation] = useState(null);

  const handleSearch = async (query) => {
    try {
      // Using Nominatim (OpenStreetMap's search engine)
      const response = await axios.get(
        `https://nominatim.openstreetmap.org/search?format=json&q=${query}`
      );
      setSearchResults(response.data);
      onSearch(response.data);
    } catch (error) {
      console.error('Search error:', error);
    }
  };

  return (
    <div className="search-add-container">
      <div className="search-box">
        <input
          type="text"
          placeholder="Search for places..."
          onChange={(e) => handleSearch(e.target.value)}
        className="search"/>
        <ul>
          {searchResults.map((result, index) => (
            <li key={index} onClick={() => setSelectedLocation(result)}>
              {result.display_name}
            </li> 
            // for lists
          ))}
        </ul>
      </div>

      
      
    </div>
  );
};

export default SearchAndAdd;