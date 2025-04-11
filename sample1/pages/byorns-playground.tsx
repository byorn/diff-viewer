import Link from "next/link";
import Image from "next/image";
const ByornsPlayground = ():JSX.Element=>{
    return (<>

        <p> Go to <Link href={'dynamic-route/woohooparam'}>dynamic routing</Link> my friend ! Notice link component doesnt do a full page refresh unless you use the anchor tag only </p>
        <br/>
        <br/>
        <Image src={"/hero-image.png"} width={500} height={300} alt={'myimage'}/>
          </>);
}
export default ByornsPlayground;