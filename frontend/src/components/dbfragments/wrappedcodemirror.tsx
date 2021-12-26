import React, { ForwardedRef } from "react";
import RCM from "react-codemirror";

export default function WrappedCodeMirror({
  editorRef,
  ...props
}: {
  editorRef: ForwardedRef<ReactCodeMirror.ReactCodeMirror>;
}) {
  return <RCM {...props} ref={editorRef} />;
}
