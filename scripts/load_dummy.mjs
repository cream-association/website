import dotenv from "dotenv";
import _ from "lodash";
import PocketBase from "pocketbase";
import fs from "fs";
import { exit } from "process";

dotenv.config();

const { POCKET_URL } = process.env;
const pb = new PocketBase(POCKET_URL);
const post_number = 200;
const dummyTags = [
  {
    name_en: "Technology",
    description_en: "Explore the latest in tech innovations.",
    name_fr: "Technologie",
    description_fr: "Explorez les dernières innovations technologiques.",
    color: "#FF5733",
  },
  {
    name_en: "Travel",
    description_en: "Discover new destinations and travel tips.",
    name_fr: "Voyage",
    description_fr:
      "Découvrez de nouvelles destinations et des astuces de voyage.",
    color: "#33FFB8",
  },
  {
    name_en: "Food",
    description_en: "Delicious recipes and culinary adventures.",
    name_fr: "Nourriture",
    description_fr: "Recettes délicieuses et aventures culinaires.",
    color: "#335EFF",
  },
  {
    name_en: "Fitness",
    description_en: "Stay fit and healthy with workout tips.",
    name_fr: "Fitness",
    description_fr:
      "Restez en forme et en bonne santé avec des conseils d'entraînement.",
    color: "#FF33C7",
  },
  {
    name_en: "Art",
    description_en: "Discover the beauty of artistic expression.",
    name_fr: "Art",
    description_fr: "Découvrez la beauté de l'expression artistique.",
    color: "#FFEB33",
  },
  {
    name_en: "Fashion",
    description_en: "Trendy styles and fashion inspiration.",
    name_fr: "Mode",
    description_fr: "Styles tendance et inspiration mode.",
    color: "#33FF8A",
  },
  {
    name_en: "Science",
    description_en: "Explore the wonders of science and discovery.",
    name_fr: "Science",
    description_fr:
      "Explorez les merveilles de la science et de la découverte.",
    color: "#A233FF",
  },
  {
    name_en: "Music",
    description_en: "Listen to the latest hits and music news.",
    name_fr: "Musique",
    description_fr: "Écoutez les derniers hits et les actualités musicales.",
    color: "#FF8C33",
  },
  {
    name_en: "Books",
    description_en: "Book reviews and literary discussions.",
    name_fr: "Livres",
    description_fr: "Critiques de livres et discussions littéraires.",
    color: "#33FFEC",
  },
  {
    name_en: "Photography",
    description_en: "Capturing moments through the lens.",
    name_fr: "Photographie",
    description_fr: "Capturer des moments à travers l'objectif.",
    color: "#FF3333",
  },
];

const dummyCollections = [
  {
    name_en: "Technology Collection",
    description_en: "A collection of technology-related articles.",
    name_fr: "Collection Technologie",
    description_fr: "Une collection d'articles liés à la technologie.",
  },
  {
    name_en: "Travel Collection",
    description_en: "A collection of travel guides and tips.",
    name_fr: "Collection Voyage",
    description_fr: "Une collection de guides de voyage et de conseils.",
  },
  {
    name_en: "Food Collection",
    description_en:
      "A collection of mouth-watering recipes and culinary adventures.",
    name_fr: "Collection Nourriture",
    description_fr:
      "Une collection de recettes délicieuses et d'aventures culinaires.",
  },
  {
    name_en: "Fitness Collection",
    description_en: "A collection of fitness tips and workout routines.",
    name_fr: "Collection Fitness",
    description_fr:
      "Une collection de conseils de fitness et de routines d'entraînement.",
  },
  {
    name_en: "Art Collection",
    description_en:
      "A collection of articles exploring various forms of artistic expression.",
    name_fr: "Collection Art",
    description_fr:
      "Une collection d'articles explorant diverses formes d'expression artistique.",
  },
  {
    name_en: "Fashion Collection",
    description_en: "A collection of fashion trends and style inspiration.",
    name_fr: "Collection Mode",
    description_fr:
      "Une collection de tendances mode et d'inspiration stylistique.",
  },
  {
    name_en: "Science Collection",
    description_en: "A collection of scientific discoveries and breakthroughs.",
    name_fr: "Collection Science",
    description_fr:
      "Une collection de découvertes scientifiques et de percées.",
  },
  {
    name_en: "Music Collection",
    description_en:
      "A collection of music news, reviews, and artist interviews.",
    name_fr: "Collection Musique",
    description_fr:
      "Une collection de nouvelles musicales, de critiques et d'interviews d'artistes.",
  },
  {
    name_en: "Books Collection",
    description_en: "A collection of book reviews and literary discussions.",
    name_fr: "Collection Livres",
    description_fr:
      "Une collection de critiques de livres et de discussions littéraires.",
  },
  {
    name_en: "Photography Collection",
    description_en:
      "A collection of stunning photographs and photography tips.",
    name_fr: "Collection Photographie",
    description_fr:
      "Une collection de superbes photographies et de conseils en photographie.",
  },
];

const img_content = fs.readFileSync("./img/robot.jpg");

pb.autoCancellation(false);

async function createTagAndCollection() {
  for (const tag of dummyTags) {
    await pb.collection("blog_tag").create(tag);
  }
  for (const col of dummyCollections) {
    await pb.collection("blog_collection").create(col);
  }

  let tags = await pb.collection("blog_tag").getFullList({
    sort: "-created",
  });
  tags = tags.map((item) => item.id);

  let collections = await pb.collection("blog_collection").getFullList({
    sort: "-created",
  });
  collections = collections.map((item) => item.id);

  return { tags, collections };
}

const { tags, collections } = await createTagAndCollection();

for (let i = 0; i <= post_number; i++) {
  let numberOfTag = Math.floor(Math.random() * tags.length);

  const data = {
    title_fr: `Article numéro ${i}`,
    slug_fr: `article-numero-${i}`,
    header_image: new File([img_content], "robot.jpg"),
    header_text_fr:
      "Bienvenue sur notre tout nouveau blog ! Nous sommes ravis de vous accueillir ici et avons hâte de partager avec vous tout un tas de contenus passionnants.",
    content_fr: `
    <h1>Bienvenue sur Notre Blog !</h1>

<p>Bienvenue sur notre tout nouveau blog ! Nous sommes ravis de vous accueillir ici et avons hâte de partager avec vous tout un tas de contenus passionnants.</p>

<h2>Présentation de Notre Dernier Produit : Le Widget 2.0</h2>

<p>Aujourd'hui est un jour spécial pour nous alors que nous dévoilons notre tout dernier produit : le Widget 2.0 ! Ce projet a été une véritable histoire d'amour pour notre équipe, et nous sommes extrêmement heureux de pouvoir enfin le partager avec le monde entier.</p>

<img src="https://placehold.co/600x400" alt="Widget 2.0">

<p>Le Widget 2.0 est bien plus qu'une simple mise à niveau ; c'est une révolution. Avec son design épuré, ses fonctionnalités améliorées et ses performances accrues, il est sûr de changer la façon dont vous travaillez.</p>

<h2>Comment Commencer</h2>

<p>Nous savons que vous êtes impatient de mettre la main sur le Widget 2.0, alors voici comment commencer :</p>

<ol>
  <li>Rendez-vous sur notre site web et téléchargez le logiciel Widget 2.0.</li>
  <li>Suivez les instructions d'installation simples pour configurer le logiciel sur votre appareil.</li>
  <li>Une fois installé, ouvrez l'application Widget 2.0 et suivez les instructions pour créer votre compte.</li>
</ol>

<h2>Exemple de Code</h2>

<p>Pour tous les développeurs là-bas, voici un exemple rapide de comment vous pouvez utiliser l'API Widget 2.0 pour personnaliser votre expérience :</p>

<pre><code class="language-rust">
use std::sync::mpsc;
use std::thread;
use std::time::Duration;

fn main() {
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        println!("Got: {}", received);
    }
}
</code></pre>

<h2>Regardez Notre Démo Produit</h2>

<div class="video-container">
  <iframe width="560" height="315" src="https://www.youtube.com/embed/PP8MuuDBkj0" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
</div>

<h2>Conclusion</h2>

<p>C'est tout pour le moment ! Assurez-vous de vous abonner à notre blog pour plus de mises à jour, de tutoriels et de conseils exclusifs.</p>

  `,
    collection: _.sample(collections),
    tag: _.sampleSize(tags, _.random(1, numberOfTag)),
    author: `${(+new Date() * Math.random()).toString(36).substring(0, 8)} ${(
      +new Date() * Math.random()
    )
      .toString(36)
      .substring(0, 8)}`,
    status: "Published",
    view_count: 0,
    seo_meta_title_fr: `Article numéro ${i}`,
    seo_meta_description_fr: `Viens lire notre nouvel article: Article ${i}`,
    seo_meta_keywords_fr: "robotic,article,dummy",
    title_en: `Article number ${i}`,
    slug_en: `article-number-${i}`,
    header_text_en:
      "Welcome to our brand new blog! We're thrilled to have you here and can't wait to share all sorts of exciting content with you.",
    content_en: `
<h1>Welcome to Our Blog!</h1>

<p>Welcome to our brand new blog! We're thrilled to have you here and can't wait to share all sorts of exciting content with you.</p>

<h2>Introducing Our Latest Product: The Widget 2.0</h2>

<p>Today is a special day for us as we unveil our newest product: the Widget 2.0! This has been a labor of love for our team, and we're beyond excited to finally share it with the world.</p>

<img src="https://placehold.co/600x400" alt="Widget 2.0">

<p>The Widget 2.0 is more than just an upgrade; it's a game-changer. With its sleek design, enhanced features, and improved performance, it's sure to revolutionize the way you work.</p>

<h2>How to Get Started</h2>

<p>We know you're eager to get your hands on the Widget 2.0, so here's how to get started:</p>

<ol>
  <li>Head over to our website and download the Widget 2.0 software.</li>
  <li>Follow the simple installation instructions to set up the software on your device.</li>
  <li>Once installed, open the Widget 2.0 application and follow the prompts to create your account.</li>
</ol>

<h2>Code Example</h2>

<p>For all you developers out there, here's a quick example of how you can use the Widget 2.0 API to customize your experience:</p>

<pre><code class="language-rust">
use std::sync::mpsc;
use std::thread;
use std::time::Duration;

fn main() {
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        println!("Got: {}", received);
    }
}
</code></pre>

<h2>Watch Our Product Demo</h2>

<div class="video-container">
  <iframe width="560" height="315" src="https://www.youtube.com/embed/PP8MuuDBkj0" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
</div>

<h2>Conclusion</h2>

<p>That's all for now, folks! Be sure to subscribe to our blog for more updates, tutorials, and insider tips.</p>
    `,
    seo_meta_title_en: `Article number ${i}`,
    seo_meta_description_en: `Come read our new article: Article number ${i}`,
    seo_meta_keywords_en: "article,robotic,dummy",
    description_fr:
      "Bienvenue sur notre tout nouveau blog ! Nous sommes ravis de vous accueillir ici et avons hâte de partager avec vous tout un tas de contenus passionnants.",
    description_en:
      "Welcome to our brand new blog! We're thrilled to have you here and can't wait to share all sorts of exciting content with you.",
  };

  try {
    await pb.collection("blog_post").create(data);
    console.log(`created ${i} articles`);
  } catch (error) {
    console.error(error);
  }
}
