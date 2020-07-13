import React from 'react';

type FeedProps = {
    posts: React.ReactChild[]
};

class Feed extends React.Component<FeedProps> {
    posts: React.ReactChild[]

    constructor(props: FeedProps) {
        super(props);
        this.posts = props.posts;
    }

    render() {
        return (
            <div>
                {this.posts}
            </div>
        );
    }
}

export default Feed;