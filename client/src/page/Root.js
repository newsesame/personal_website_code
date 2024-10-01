import React, { useEffect, useRef } from 'react';
import { FaLinkedin, FaGithub } from 'react-icons/fa';
import {Col, Row, Image } from "react-bootstrap";
import selfie from '../image/selfie.jpg';
import { Helmet } from 'react-helmet-async';
import cuhkicon from '../image/cuhkicon.png';


const categoryColors = {
    "Language": "hsl(208,54%, 73%)", // Deep Blue
    // "Frontend": "#b2b2b2", // 
    "Backend": "hsl(208,54%, 92%)",  // Pale Blue
    "Database": "hsl(208,54%, 73%)",
    "Others": "hsl(208,54%, 92%)",  
    
};


const techStack = {
    // Images stored in public
    "Language": [
        { name: 'Python', image: 'Python.png', familiarity: "Medium" },
        { name: 'GO Lang', image: 'Golang.png',familiarity: "Medium" },
        { name: 'Java', image: 'Java.png',familiarity: "Medium" },
    ],

    "Backend": [
        { name: 'Node.js', image: 'nodejs.png', familiarity: "Medium" },
        { name: 'Golang', image: 'gofiber.png', familiarity: "Medium"},
    ],
    "Database": [
        { name: 'MongoDB', image: 'mongodb.svg', familiarity: "Medium" },
        { name: 'Postgresql', image: 'postgresql.png', familiarity: "Basic" },
        
    ],
    "Others": [
        { name: 'Selenium', image: 'selenium.webp', familiarity: "Medium" },
        { name: 'Docker', image: 'Docker.png', familiarity: "Medium" },
        
    ],
};


const Root = () => {


    // Typing Effect
    const divRef = useRef(null);
    const text = "(I guess I am really bad at Frontend development.... I feel so sorry if the design is unpleasant to you. -,-)";

    let typingTimer = null; // Counter
    
    function textTypingEffect(element, text, i = 0) {
        // Reset the content at the start
        if (i === 0) {
            element.textContent = "";
        }
        
        // add every wording everytime
        element.textContent += text[i]; 
    
        // End Condition
        if (i < text.length - 1) {
            typingTimer = setTimeout(() => textTypingEffect(element, text, i + 1), 70);
        }
    }
    
    useEffect(() => {
        if (divRef.current) {
            // Make sure only one counter is working, for the stability 
            if (typingTimer) {
                clearTimeout(typingTimer);
            }
    
            textTypingEffect(divRef.current, text); 
        }
    
        return () => {
            if (typingTimer) {
                clearTimeout(typingTimer);
            }
        };
    }, [text,textTypingEffect, typingTimer]);
    

    return (
        <>
        <Helmet>
        <title>Josh Chau</title>
        <meta name="description" content="Josh Chau personal website home page." />
        </Helmet>
        <div fluid>

            <div  className='Container text-center'>
                <h1 style={{ flex: '1 1 100px' }}>Welcome to My Personal Website</h1>
            </div>
            
            <div  className='Container' style={{justifyContent: 'center'}} >
                <p ref={divRef} style={{ fontSize: "17px" }}>
                </p>
            </div>

            
            <div  className='Container text-center space' >

                <p style={{ flex: '1 1 100px' }}className='text-center'> 
                This website is built around node.js, Go Lang and MongoDB.<br></br> I made this website in order to train my full stack skills. You may check out the source code on my github.
                <br></br>
                It showcases my education and projects. I will also share some personal things and ideas here. 🙂 
                </p>
            </div>


            <div className='Container space'>
                <div style={{ flex: '1 1',}}>
                    <Image className="selfie items" src={selfie} roundedCircle />
                </div>
                
                <div className="items" style={{ flex: '3 1' }}>
                    <h2>About Me</h2>
                    <p>
                        Hello. I am a Year 3 Computer Science student at the Chinese University of Hong Kong, with a minor in Statistics. Currently, I'm pursuing my minor in statistics and 
                        working on various software development projects.
                        
                    </p>
                    <Col className='no-padding'>

                        <Row className="align-items-center">
                            <Col> Please visit my Project on Github:<a href="https://github.com/newsesame" target="_blank" rel="noopener noreferrer"><FaGithub size={50} color="black"/></a></Col>
                  
                                
                        </Row>

                        <Row className="align-items-center">
                    
                            <Col>Also, feel free to connect me via linkedin:<a href="https://www.linkedin.com/in/josh-chau/" target="_blank" rel="noopener noreferrer"><FaLinkedin size={50} style={{ marginRight: '10px' }} /></a></Col>
        
                        </Row>

                        <Row className="align-items-center">
                            <Col>
                                I do Leetcode sometimes as well! 
                                <a href="https://leetcode.com/u/newsesame/" target="_blank" rel="noopener noreferrer">
                                <Image  style={{ 

                                        marginTop:"7.5px",
                                        marginBottom:"7.5px"
                                        
                                    }} src="Leetcode.png" alt="photo"   height="35px" roundedCircle/>
                                </a>
                            </Col>
                  
                                
                        </Row>
                    </Col>

                    
                    
                </div>

            </div>
            
            <div style={{marginBottom: "0px",}} className='Container space'>

                <h2>Education</h2>

            </div>


            <div className='Container space'>
                <div style={{ flex: '1 1' }}><Image height="" src={cuhkicon} rounded className='shadow cuhkicon'/></div>

                <div style={{ flex: '5 1', paddingLeft: "40px" }} >
                    <h3 >B.Sc. in Computer Science <br></br>Minor in Statistics</h3>
                    <p>The Chinese University of Hong Kong<br></br>Sept 2022 to Present</p>
                </div>
            </div>

           
            <div style={{marginBottom: "0px",}} className='Container space'>
                <h2>Tech Stack</h2>
            </div>


            <div className='Container'>

                <Col>
                    {Object.keys(techStack).map((category, index) => (
                    
                    <Row 
                        key={index}                            
                        style={{ 
                        backgroundColor: categoryColors[category],
                        borderRadius: "15px", // 圓角效果,
                        marginBottom: "10px",
                        padding: "15px"
                        }}>

                        {/* Title Column */}
                        <Col md="3">
                        <h3 style={{marginBottom: "0px",}}>{category}</h3>
                        <p>(Familiarity)</p>
                        </Col>

                        {/* Element Columns */}
                        {techStack[category].map((tech, idx) => (
                            <Col 
                                key={idx} 
                                className="d-flex flex-column align-items-center" >
                                <Row>
                                    <Image  
                                        style={{ 
                                            backgroundColor: "#fff", 
                                            borderRadius: "50%", 
                                            padding: "15px",           
                                            boxShadow: "0 4px 8px rgba(0, 0, 0, 0.2)", 
                                            marginBottom: "5px",
                                            width: "90px",  
                                            height: "90px", 
                                            objectFit: "cover" 
                                        }} 
                                        className="shadow" 
                                        src={tech.image} 
                                        alt={tech.name} 
                                    />
                                </Row>
                                
                                <Row>
                                    <p style={{ marginBottom: "0px" }}>{tech.name}</p>
                                </Row>
                            </Col>
                        ))}


                    </Row>
                        
                    ))}
                </Col>
                
            </div>

        </div>
        </>
    )
}

export default Root