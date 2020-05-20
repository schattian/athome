import React from 'react';
import {
  Danger,
  Info,
  Muted,
  Primary,
  Quote,
  Success,
  Warning,
} from 'src/components/Typography';


const TypographyExample = () => (
  <div>
    <Quote
      text="I will be the leader of a company that ends up being worth billions of dollars, because I got the answers. I understand culture. I am the nucleus. I think that’s a responsibility that I have, to push possibilities, to show people, this is the level that things could be at."
      author=" Kanye West, Musician"
    />
    <Muted>
      Muted Text
    </Muted>
    <Primary>
      Primary Text
    </Primary>
    <Info>
      Info Text
    </Info>
    <Success>
      Success Text
    </Success>
    <Warning>
      Warning Text
    </Warning>
    <Danger>
      Danger Text
    </Danger>
  </div>
);

export default TypographyExample;
