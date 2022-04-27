import axios from 'axios';
import adapter from "axios/lib/adapters/http"

axios.defaults.adapter = adapter;

export class VideosAPI {
    constructor(url) {
        this.url = url;
    }

    async getVideoList() {
        return axios.get(this.url)
            .then(data => data.data);
    }
}

export default new VideosAPI("https://my-json-server.typicode.com/modanisa/bootcamp-video-db/videos");