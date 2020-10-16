import React from 'react';
import './App.css';
import './tailwind.output.css';
import logo from './images/logo.svg';
import Post from './components/Post';
import Feed from './components/Feed';

function App() {

  const post = <Post content="My first fake twitter post" pictureFile={logo}></Post>;

  const posts = [post, post, post];

  return (
    <div className="mx-auto flex p-6 bg-white rounded-lg shadow-xl">
      <Feed posts={posts} ></Feed>
    </div>
  );
}

export default App;
