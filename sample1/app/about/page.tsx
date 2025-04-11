const AboutPage = async() => {
    const dataForPage = await getData();
    return <>
        {JSON.stringify(dataForPage)}
    Hello World</>
}
export default AboutPage;

const getData = async ():Promise<any> => {
    console.log("runs on the server");
    const response =  await fetch('https://dummyjson.com/comments', { next: { revalidate: 10}});
    return response.json();
}