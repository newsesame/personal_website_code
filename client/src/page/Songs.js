import React from 'react';
import axios from 'axios';


import {Col, Row} from "react-bootstrap";
import { useEffect, useState } from 'react';

import 'bootstrap/dist/css/bootstrap.min.css';





// import Education from '../page/Education';
const Root = () => {

    useEffect(() => {
		document.title = "Playlist — Josh Chau"
	}, []);
    // Get the song data
    const [songs, setSongs] = useState([]);
    const apiUrl = process.env.REACT_APP_WEBSERVER_API_ROOT;
    console.log(apiUrl);
    useEffect(() => {
      const fetchSongs = async () => {
        try {
            const response = await axios.get(apiUrl+ "/songs");
            setSongs(response?.data?.song_records);
        } catch (error) {
            console.error('Error fetching songs:', error);
        }
      };
  
      fetchSongs();
    }, [apiUrl]);

    console.log(songs);

    return (
        <>
            <div className='Container-h'>
                <h2>My Playlist</h2>
                <p className="no-margin">All song information is captured by the Python + Selenium web scraper script and stored in the MongoDB.</p>
                <hr></hr>
    
                {songs.map((group, index) => {
                    
                    const shouldRenderYear = index === 0 || group.year !== songs[index - 1].year;

                    // Every Year
                    return (
                        <div key={index}>
                            
                            {shouldRenderYear && <h2>{group.year}</h2>}
                            <div
                                style={{
                                    // backgroundColor: "#DCDCDC",
                                    // borderColor: "#343a40",
                                    // color: "#343a40",
                                    padding: "0 px",
                                    width:"360px"}}>
                                <button
                                type="button"
                                data-bs-toggle="collapse"
                                data-bs-target={`#collapse-${index}`}
                                aria-expanded="true"
                                aria-controls={`#collapse-${index}`}
                                className="styled-button"
                                >
                                <p className="button-text">{group.month}</p>
                                </button>

                            </div>
                            <div
                                id={`collapse-${index}`}
                                className="accordion-collapse collapse show"
                                data-bs-parent="#accordionExample">
                                <p> ⬆️⬆️ Click the month number above to expand/collapse the playlist!</p>
                                <div style={{ border: '1px solid black', padding: "20px" }}>
                                <p>List: </p>

                                <ul className='p-0'>
                                {group.songs.map((song) => (
                                    <li key={song.song_id}>
                                            
                                            <Row>
                                                <Col md={3}>
                                                    <div className='cover'>
                                                        <img src={apiUrl+"/cover/" + song.cover_image} alt="song cover" />
                                                    </div>
                                                </Col>

                                                <Col md={4} className="d-flex flex-column justify-content-center">
                                                    <div style={{paddingLeft:"30px"}}>
                                                    <Row><h4 className='p-0'>{song.title}</h4></Row>
                                                    <Row>{song.artist}</Row>
                                                    <Row>{song.album}</Row>
                                                    </div>
                                                </Col>
                                                <Col md={3} className="d-flex flex-column justify-content-center">
                                                <Row dangerouslySetInnerHTML={{ __html: song.lyricist ? song.lyricist.replace(/\n/g, '<br>') : "Unknown" }} />
                                                    <Row>{song.composer}</Row>
                                                </Col>
                                                <Col md={2} className="d-flex flex-column justify-content-center">
                                                <a href={song.song_url} target="_blank" rel="noopener noreferrer">
                                                    <button
                                                        style={{
                                                        backgroundColor: "hsl(208,54%, 73%)",  
                                                        color: "white",              
                                                        padding: "5px 10px",        
                                                        border: "none",              
                                                        borderRadius: "8px",         
                                                        boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)", 
                                                        cursor: "pointer",           
                                                        transition: "background-color 1.0s ease", 
                                                        }}
                                                        onMouseEnter={(e) => {
                                                        e.target.style.backgroundColor = "hsl(208,54%, 92%)"; 
                                                        }}
                                                        onMouseLeave={(e) => {
                                                        e.target.style.backgroundColor = "hsl(208,54%, 73%)"; 
                                                        }}
                                                    >
                                                        KKBOX
                                                    </button>
                                                    </a>
                                                </Col>
                                            </Row>
                                        {/* </div> */}
                                    </li>
                                ))}                                    

                                </ul>
                                </div>
                            </div>
                        </div>
                    );
                })}
            </div>
        </>
    );
    
}

export default Root