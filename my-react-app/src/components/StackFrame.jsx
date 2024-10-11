
const StackFrame = ({number, id}) => {

    return (
        <>
            <h1 id={"stackframe" + id} style={{border: "solid red", width:"100px", height:"50px"}}>{number}</h1>
        </>
    )

}

export default StackFrame