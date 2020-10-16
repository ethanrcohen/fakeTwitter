import React from 'react';

type FeedProps = {
    posts: React.ReactChild[]
};

const Feed = (props: FeedProps) => {
    return (
        <div>
            {props.posts}
        </div>
    );
};

export default Feed;