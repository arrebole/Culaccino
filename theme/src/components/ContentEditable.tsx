import React from 'react';

interface Props {
    handleChange(e: any, target: string): void
    content: string
}

export default class ContentEditable extends React.Component<Props> {
    constructor(props: Props) {
        super(props);
        this._onChange = this._onChange.bind(this);
    }
    
    private elem: HTMLDivElement | null = null

    _onChange(ev: any) {
        const value = this.elem != null ? this.elem['innerText'] : "";
        this.props.handleChange(value, "content");
    }

    render() {
        return (
            <div
                style={{minHeight:"500px", width:"100%", maxWidth:"980px", boxSizing:"border-box",whiteSpace: 'pre-line'}}
                className="bg-white border p-2"
                ref={elem => this.elem = elem}
                dangerouslySetInnerHTML={{ __html: this.props.content }}
                onBlur={this._onChange}
                contentEditable
            />
        )
    }
}
