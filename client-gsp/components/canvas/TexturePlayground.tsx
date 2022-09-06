import React from 'react';
import styles from './playground.module.css';
import PropTypes from 'prop-types';
import { Stage, Layer, Line } from 'react-konva';
import { KonvaEventObject } from 'konva/lib/Node';
import Button from '../Button';

const DrawArea = React.forwardRef((props, ref)=> {
    const [lines, setLines] = React.useState<Array<{points: Array<number>, tool: String}>>([]);
    const isDrawing = React.useRef(false);

    React.useImperativeHandle(ref, () => ({

        clearLines() {
            setLines([]);
        },

        getLines() {
            return lines
        }
    
    }));
    
    const handleMouseDown = (e: KonvaEventObject<MouseEvent>) => {
        isDrawing.current = true;
        const evt = e.target.getStage()
        if (evt) {
            const pos = evt.getPointerPosition();
            if (pos) {
                setLines([...lines, { points: [pos.x, pos.y], tool: "" }]);
            }
            
        }
        
    };
    
    const handleMouseMove = (e: KonvaEventObject<MouseEvent>) => {
        // no drawing - skipping
        if (!isDrawing.current) {
          return;
        }
        const stage = e.target.getStage();
        if (!stage) {
            return
        }
        const point = stage.getPointerPosition();
        if (!point) {
            return
        }
    
        // To draw line
        let lastLine = lines[lines.length - 1];
        
        if(lastLine) {
            // add point
            lastLine.points = lastLine.points.concat([point.x, point.y]);
                
            // replace last
            lines.splice(lines.length - 1, 1, lastLine);
            setLines(lines.concat());
        }
        
    };
    
    const handleMouseUp = () => {
        isDrawing.current = false;
    };

    return (
        <div className=" text-center text-dark">
            <Stage
                width={300}
                height={300}
                onMouseDown={handleMouseDown}
                onMousemove={handleMouseMove}
                onMouseup={handleMouseUp}
                className="canvas-stage"
            >
                <Layer>
                    {lines.map((line, i) => (
                        <Line
                        key={i}
                        points={line.points}
                        stroke="#df4b26"
                        strokeWidth={2}
                        tension={0.5}
                        lineCap="round"
                        globalCompositeOperation={
                            line.tool === 'eraser' ? 'destination-out' : 'source-over'
                        }
                        />
                    ))}
                </Layer>
            </Stage>
        </div>
    )
})

type PlaygroundProps = {
    onSave: React.MouseEventHandler<HTMLDivElement>
};

const TexturePlayground = function({onSave}: PlaygroundProps) {
    const drawingRef = React.useRef();

    const handleEraseButton = () => {
        if (drawingRef.current) {
            drawingRef.current.clearLines() // clear canvas
        }
        
    }

    const handleSaveButton = () => {
        if (drawingRef.current) {
            const lines = drawingRef.current.getLines()
            onSave(lines)
        }
    }
    
    return (
        <React.Fragment>
            <div className={styles['helper-title']}>
                Draw your login signature! This can be whatever you find that represent yourself.
            </div>
            <div className={styles['action-row']}>
                <Button onClick={handleEraseButton}>Erase</Button>
                <Button onClick={() => handleSaveButton()}>Save</Button>
            </div>
            
            <div className={styles['area']}>
                <DrawArea ref={drawingRef}/>
             </div>
        </React.Fragment>
    )
}

TexturePlayground.propTypes = {
    data: PropTypes.any
}

export default TexturePlayground;