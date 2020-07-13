import React from 'react';
import './App.css';
import Post from './components/Post';
import Feed from './components/Feed';

function App() {

  const post = <Post content="My first fake twitter post"></Post>;

  const posts = [post, post, post];

  return (
    <div className="App">
      <Feed posts={posts} ></Feed>
    </div>
  );
}

export default App;
