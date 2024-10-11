import useFetch from "react-fetch-hook";
import StackFrame from "./StackFrame";



const Stack = ({}) => {
    const {isLoading, data, error} = useFetch(
        "https://stackmachine.com/stack"
    )

    if (error){
        return <p> Error loading the stack.</p>
    }

    // TODO handle isLoading
    // if (isLoading){

    // }

    if (data) {
        return (
            <div>
                {data.array.map((frame, _) => (
                    <StackFrame number={frame.stackNumber} id={frame.stackID}></StackFrame>
                ))}
            </div>
        );
    }

}

export default Stack;