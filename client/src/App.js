// import logo from './logo.svg';

import './App.css';
import './Large.css';
import {Routes, Route} from "react-router-dom";
import Header from "./components/layout/Header";
import { HelmetProvider } from 'react-helmet-async';

import Root from './page/Root';
import Test from './page/Test';
import Songs from "./page/Songs";
import Career from "./components/Career";

import { Col, Row } from 'react-bootstrap';




function App() {

  return (
	<HelmetProvider>
    <>
		
		<div style={{ paddingTop: '70px' }}></div>

		{/* Header */}
		<Row>
			<Col md="6"className="mx-auto" >
				<Header/>
			</Col>
		</Row>

		{/* Routing */}
		<Routes >
		<Route path = "/" element={<Root/>}/>
		<Route path = "/test" element={<Test/>}/>
		<Route path = "/songs" element={<Songs/>}/>
		<Route path = "/career" element={<Career/>}/>
		 
			{/* <Route path = "/contacts" element={<Contacts/>}/> */}


		</Routes>
    </>
	</HelmetProvider>
  );
}

export default App;
