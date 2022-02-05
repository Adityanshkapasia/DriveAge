import React from 'react';
import { homeObjOne, homeObjTwo, homeObjThree, homeObjFour } from './Data';
import { InfoSection, Pricing } from '../../components';
import CarouselSlider from '../../components/CarouselSlider';
import { CarouselData } from '../../components/CarouselData';


function Home() {
  return (
    <>
      <InfoSection {...homeObjOne} />
      {/*  <InfoSection {...homeObjThree} />*/}
      <div>
        <CarouselSlider slides={CarouselData}/>
      </div>
      <InfoSection {...homeObjTwo} />
      <Pricing />
      <InfoSection {...homeObjFour} />
    </>
  );
}

export default Home;
