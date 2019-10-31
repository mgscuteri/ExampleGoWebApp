import React, { Component } from 'react';

class HomePage extends React.Component {  
    constructor(props) {
      super(props);
    }

    render() {
      return (
        <div>
          <h2>
            Welcome to the Mike Scuteri's PythonPoweredPi!
          </h2>
          <p>
            The year is 2019.  The world has all but given up on owning your own hardware.  “Hosting a website on your own hardware is too expensive!” they say.  In stark contrast to that sentiment, I humbly present the PythonPoweredPi, a webapp that utilizes some of the most modern web technologies, and is being 100% hosted on $35 worth of hardware.   
            
          </p>
          <p>
            The website is still a work in progress.  Below, is an arbitrary table that will write records to disk.  Below that, is a summary of the technologies being used.
          </p>
          <br/>
        </div>
      );
    }
  }

export default HomePage
