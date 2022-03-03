import React from 'react';

import './home.scss';
import { Link } from 'react-router-dom';

const Home = () => (
  <div className="home-container">
    <div className="hero-container">
      <div>
        <h1>Brentwood, NJ</h1>
        <p>Rated the safest town within 50 miles!</p>
      </div>
      <div className="scroll-down">
        <span>Scroll Down</span>
        <i
          className="fa fa-chevron-down fa-bounce"
          style={{
            '--fa-bounce-start-scale-x': '1', '--fa-bounce-start-scale-y': '1', '--fa-bounce-jump-scale-x': '1', '--fa-bounce-jump-scale-y': '1', '--fa-bounce-land-scale-x': '1', '--fa-bounce-land-scale-y': '1', '--fa-bounce-rebound': '0', margin: '0.5rem',
          }}
        />
      </div>
    </div>
    <div className="latest-news">
      <h1>Latest Brentwood News</h1>
      <div className="news-item-list">
        <div className="news-item">
          <img src="https://via.placeholder.com/300x150" alt="headline" />
          <h3 className="headline">Lorem Ipsum</h3>
          <p className="description">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
        </div>
        <div className="news-item">
          <img src="https://via.placeholder.com/300x150" alt="headline" />
          <h3 className="headline">Lorem Ipsum</h3>
          <p className="description">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
        </div>
        <div className="news-item">
          <img src="https://via.placeholder.com/300x150" alt="headline" />
          <h3 className="headline">Lorem Ipsum</h3>
          <p className="description">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
        </div>
        <div className="news-item">
          <img src="https://via.placeholder.com/300x150" alt="headline" />
          <h3 className="headline">Lorem Ipsum</h3>
          <p className="description">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
        </div>
      </div>
    </div>
    <div className="call-to-action">
      <div className="title">
        <h2>Need COVID-19 Assistance?</h2>
        <p><em>Call our COVID-19 Hotline</em></p>
      </div>
      <div className="action">
        <a className="btn" href="#">Call Us!</a>
      </div>
    </div>
    <div className="newsletter-container">
      <a className="btn" href="#">Read More</a>
      <div className="title">
        <h1>Newsletter</h1>
        <p>Come read our monthly newsletter!</p>
      </div>
    </div>
    <div className="footer">
      <div className="contact-details">
        123 Main St.
        <br />
        Brentwood, NJ 12345
      </div>
      <div className="copyright">
        &copy; 2022 Made up design agency
      </div>
      <div className="links">
        <a href="#">About Us</a>
        <a href="#">Town Board</a>
        <Link to="/drop">Submit Feedback</Link>
      </div>
    </div>
  </div>
);

export default Home;
