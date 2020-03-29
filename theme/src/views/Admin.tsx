import React from 'react'
import Header from '../components/Header'
import api, { Paper } from '../api'
import { Link } from "react-router-dom"

const Table: React.FC<{ data: Paper[] }> = (props) => {
    let list = props.data.map(item => {
        return (
            <li key={item.title} className="d-grid grid-col-10 f5 p-2">
                <div>Title: {item.title}</div>
                <div>Type: {item.type}</div>
                <div>Create: {item.create_at}</div>
                <div>Update: {item.create_at}</div>
                <div><Link to={`/editor?title=${item.title}&tab=update`}>修改</Link></div>
            </li>
        )
    })

    return (
        <ul>{list}</ul>
    )
}


export default class AdminPage extends React.Component<{}, { papers: Paper[] }>{
    constructor(props: {}) {
        super(props)
        this.state = {
            papers: []
        }
    }

    componentDidMount() {
        api.fetchPapers().then(res => {
            this.setState({ papers: res })
        })
    }

    render() {
        return (
            <div id="admin" className="bg-gray-light">
                <Header />
                <main>
                    <div className="d-flex flex-recolumn p-2 m-2">
                        <Link to="/editor?tab=new">
                            <button className="btn-primary f5 p-2 m-1 rounded-2 border">New Paper</button>
                        </Link>
                    </div>
                    <Table data={this.state.papers} ></Table>


                </main>
            </div>
        )
    }
}