import React, { useEffect } from 'react'

const Career = () => {

    useEffect(() => {
		document.title = "Career — Josh Chau"
	}, []);


    return (
        
        <>
        <div className='Container' style={{marginBottom:"0px"}}>
            <h2 style={{ flex: '1 1 100px' }}>Career<hr></hr></h2>
                
        </div>
        
        <div className="Container">
            <p>I would like to become an Quant Developer.
                <br></br>
                <br></br>
            "In the whole branch of computer science, programming is not even 10." This is one of the things that my professor said that impressed me a lot.
            I've always believed that as a computer science undergraduate, my value is not in the fact that I can just call different libraries/technology piles. I also don't want this to be my career for the next few decades, even though I understand that we are required to be familiar with a lot of applications on the job. 
            <br></br>
            <br></br>
            But when I ask myself what I know in theory, I realise I don't know much lol... I have been thinking about what advantages I have over university students majoring in other disciplines.
            Algorithm? Data Structure? Automata Theory? And what are my advantages over other computer science students?

            <br></br>
            <br></br>
            Therefore, I started on the path of pursuing a minor in Statistics. Looking forward to one day I can combine my knowledge of computers and statistics to create something great, but not just by importing models from Keras/ Tensorflow and fitting them into data. 🥲
            <br></br>
            <br></br>
            </p>
        </div>
        </>
    
    )
}

export default Career