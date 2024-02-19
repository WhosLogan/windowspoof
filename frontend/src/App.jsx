import {Text} from "./components/text.jsx";
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from "./components/table.jsx";
import {Input} from "./components/input.jsx";
import {Button} from "./components/button.jsx";
import {BeginFetchingWindows} from "../wailsjs/go/main/App.js";
import {useEffect, useState} from "react";
import {EventsOn} from "../wailsjs/runtime/runtime.js";

function App() {
    const [windows, setWindows] = useState([]);

    useEffect(() => {
        BeginFetchingWindows();
    }, []);

    EventsOn('window-update', (w) => {
        setWindows(w)
    })

    const shortenString = (str) => {
        return str.length < 15 ? str : `${str.substring(0, 15)}...`
    }

    return (
        <div id="App">
            <div className="flex draggable justify-center pt-2 pb-3">
                <Text className="select-none">Window Spoof</Text>
            </div>
            <div className="mx-8 border-white/15 border rounded-lg px-3 pt-2 mt-2">
                <Table className="max-h-[37rem]">
                    <TableHead>
                        <TableRow>
                            <TableHeader>Process ID</TableHeader>
                            <TableHeader>Name</TableHeader>
                            <TableHeader>Window Title</TableHeader>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {windows.map(w => (
                            <TableRow key={w.pid}>
                                <TableCell className="text-zinc-200">{w.pid}</TableCell>
                                <TableCell className="text-zinc-200">{shortenString(w.name)}</TableCell>
                                <TableCell className="text-zinc-200">{shortenString(w.windowName)}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </div>
            <div className="fixed bottom-0 w-full">
                <div className="flex justify-center gap-x-6 items-center p-4 border rounded-lg border-white/15 mb-6 mx-10">
                    <Input className="max-w-48" placeholder="New Title"/>
                    <Button className="max-h-10">Change Title</Button>
                </div>
            </div>
        </div>
    )
}

export default App
