import React, { Component } from 'react'
import axios from 'axios';
class GetCaptionComponent extends Component {

    constructor() {
        super();
        this.state = {
            getcaption: 'pending'
        }
    }

    componentWillMount() {
        axios.get('api/v1/getcaption')
            .then((response) => {
                this.setState(() => {
                    return { getcaption: response.data }
                })
            })
            .catch(function (error) {
                console.log(error);
            });

    }

    render() {
        return <h1>GetCaptionComponent {this.state.getcaption}</h1>;
    }
}

export default GetCaptionComponent;
