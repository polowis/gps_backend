import React from 'react';
import styles from './playground.module.css';
import PropTypes from 'prop-types';

interface ImageLoad {
    img: HTMLImageElement;
    loaded: boolean;
}

const TexturePlayground = function({data}: any) {
    const canvasRef = React.useRef<HTMLCanvasElement>(null);
    const boardWidth = 300;
    const boardHeight = 250;
    const images: {
        [id: string]: ImageLoad
    } = {};

    /**
     * 
     * @param {*} key - the key of the image to be accessed later on, duplicate key will override
     * @param {*} filePath - the path to image file can be remote url
     */
    const loadImage = (key: string, filePath: string) => {
        const img = new window.Image();
        images[key] = {
            "img": img,
            "loaded": false
        }
        img.addEventListener("load", function () {
            images[key].img = img;
            images[key].loaded = true;
        });
        img.setAttribute("src", filePath)
    }

    /**
     * Get image by key
     * @param {*} key 
     */
    const getImage = (key: string) => {
        return images[key].img;
    }

    const setup = () => {
        data.data.images.forEach((item: any) => {
            loadImage(item.id, item.url)
        })
    }

    const drawImages = (ctx: CanvasRenderingContext2D) => {
        let x: number = 0;
        let scale = 64; // 5 image per row, 16px * 5 = 320px, 
        //canvas fixed width = 400px - padding 16 * 2 = 368
        let padding = (368 - 320) / 5
        for(const [key, obj] of Object.entries(images)) {
            ctx.drawImage(images[key].img, x, 0, 64, 64)
            x += scale + padding;
        }
    }

    React.useEffect(() => {
        setup();
        if (canvasRef.current) {
            const canvas: HTMLCanvasElement = canvasRef.current
            const context: CanvasRenderingContext2D|null = canvas.getContext('2d')

            if (!context) {
                return
            }

            //Our first draw
            canvas.style.width ='100%';
            canvas.style.height='100%';
            // ...then set the internal size to match
            canvas.width  = canvas.offsetWidth;
            canvas.height = canvas.offsetHeight + 100;

            const render = () => {
                context.fillStyle = '#000000'
                context.fillRect(0, 0, context.canvas.width, context.canvas.height);
                let loaded = true;
                for(const [key, obj] of Object.entries(images)) {
                    if (!obj.loaded) {
                        loaded = false
                        break
                    }
                }
                if (loaded) {
                    drawImages(context)
                }
                requestAnimationFrame(render);
            }
            render();
        }
       
        
    })

    return (
        <div className={styles['canvas']}>
            <canvas ref={canvasRef}/>
        </div>
    )
}

TexturePlayground.propTypes = {
    data: PropTypes.any
}

export default TexturePlayground;