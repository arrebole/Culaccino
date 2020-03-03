import React from 'react';
import Header from '../components/Header'
import Unit from '../components/Unit'
import api, { Paper } from "../api"


class PaperList extends React.Component<{},{ papers: Paper[] }> {
    constructor(props: {}){
        super(props)
        this.state = {
            papers: []
        }
    }
    componentDidMount(){
        api.fetchPapers().then(res=>{
            this.setState({ papers: res.data})
        })
    }
    render() {
        return (
            this.state.papers.map(item =>
                <article className="m-3" style={{ width: "95%", maxWidth: "900px" }} key={item.title}>
                    <Unit {...item}></Unit>
                </article>
            )
        )
    }

}
export default class HomePage extends React.Component {
    render() {
        return (
            <div id="home" className="bg-gray-light">
                <Header />
                <main className="d-flex flex-column flex-items-center">
                    <PaperList />
                </main>
            </div>
        )
    }
}