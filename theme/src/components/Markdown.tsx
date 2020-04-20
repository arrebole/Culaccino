import React from "react";
import marked from 'marked';
import hljs from 'highlight.js'

export default function Markdown(props: { contents: string }) {
    marked.setOptions({ highlight: (code, lang) => hljs.highlight(lang, code).value });
    const md = { __html: marked(props.contents) };
    return (
        <section className="bg-white border border-gray-400 shadow rounded-sm" >
            <div className="text-base p-8 markdown-body" dangerouslySetInnerHTML={md}></div>
        </section>
    )
}
