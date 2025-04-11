import { useRouter } from 'next/router'
import Head from "next/head";
import jd from "../../data/jobs.json";
import {GetStaticProps, GetStaticPaths} from "next";
import JobDataType from "../types/JobDataType";


interface Props {
    data: JobDataType
}


const JobDetailsPage = (props: Props)=>{
    const router = useRouter();
    return (<><Head> <title>{router.query.id}</title></Head>
        Job Details of :
        <div>
            <h1>{props.data.jobDescription}</h1>
        </div>
    </>);
}

export const getStaticProps: GetStaticProps<Props> = async (context) => {
        const jobData:JobDataType[] = jd;
        const { params } = context;
        const foundData: JobDataType | undefined = jobData.find((jobData:JobDataType) => {
                    return jobData.jobId.toString() === params?.id; //dynamic id
                    });
        const renderJobData = foundData || {
                jobId: 0,
                jobDescription: 'not found',
                jobImage: 'null',
         }

        return {
            props: { data: renderJobData }
        };
}



export async function getStaticPaths() {
    const jobData:JobDataType[] = jd;
    const paths = jobData.map((jobDataType) => {
        return {
            params: {
                id: jobDataType.jobId.toString(),
            },
        };
    });
    return {
        paths,
        fallback: true,
    };
}

export default JobDetailsPage;