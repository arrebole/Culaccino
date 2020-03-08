import React from "react"

interface Props {
    handleChange(event: any, target?:string): void;
    value:string;
}

export default class EditorPage extends React.Component<Props> {
    constructor(props: Props) {
        super(props)
        this.emitChange = this.emitChange.bind(this)
    }

    emitChange(event: any) {
        this.props.handleChange(event, "content")
    }

    render() {
        return (
            <div
                id="contentEditable"
                contentEditable
                onBlur={this.emitChange}
                dangerouslySetInnerHTML={{__html: this.props.value}}
                className="bg-white border p-3"
                style={{ minHeight: '530px', width: '100%', maxWidth: '980px', boxSizing: 'border-box',whiteSpace: 'pre-line' }}>
            </div>
        )
    }
}