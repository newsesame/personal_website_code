import axios from 'axios';
import React from 'react';


import {Container, Row, } from "react-bootstrap";
import { useEffect, useState } from 'react';


const Test= () => {

  useEffect(() => {
		document.title = "Test"
	}, []);
    const [songs, setSongs] = useState([]);
    const test = 5;
    const apiUrl = process.env.REACT_APP_WEBSERVER_API_ROOT;
    console.log(apiUrl);
    useEffect(() => {
      const fetchSongs = async () => {
        try {
            // const apiUrl = process.env.REACT_APP_WEBSERVER_API_ROOT;
            const response = await axios.get(apiUrl+ "/songs");
            setSongs(response?.data?.song_records);
        } catch (error) {
          console.error('Error fetching songs:', error);
        }
      };
  
      fetchSongs();
    }, [apiUrl]);
    console.log(test);
    console.log(apiUrl);
    console.log(songs);


    return (
        <Container>
        <Row>
            <p>Hi</p>
            
        </Row>
      </Container>
    )
}

export default Test