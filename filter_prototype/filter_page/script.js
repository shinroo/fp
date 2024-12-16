const pets = [
  {
    "name": "Blackie",
    "image": "https://spcadbn.org.za/petimages/2024-12-06_blackie.jpg",
    "description": "Blackie is a lovable girl with a calm and well behaved nature. She is a 1 year old Labrador cross. Unwanted. Kennel B14 Ref: 121763.",
    "sex": "female",
    "age": 1.0,
    "kennel": "B14",
    "reference": "121763",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Malley",
    "image": "https://spcadbn.org.za/petimages/2024-11-29_B13malley.jpg",
    "description": "Malley is a playful little girl and loves to play with toys. She is a 10 month old Labrador cross. Stray. Kennel B12 Ref: 121092.",
    "sex": "female",
    "age": 0.8333333333333334,
    "kennel": "B12",
    "reference": "121092",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Rex",
    "image": "https://spcadbn.org.za/petimages/2024-11-29_B10rex.jpg",
    "description": "Rex is a well behaved boy with a calm and chilled nature. He is a 3 year old Boerboel cross. Stray. Kennel B10 Ref: 121178.",
    "sex": "male",
    "age": 3.0,
    "kennel": "B10",
    "reference": "121178",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Lizzy",
    "image": "https://spcadbn.org.za/petimages/2024-11-29_B9lizzy.jpg",
    "description": "Lizzy is a sweet young girl with a lovable and friendly nature. She is a 9 month old Staffie cross Labrador. Unwanted. Kennel B9 Ref: 121708.",
    "sex": "female",
    "age": 0.75,
    "kennel": "B9",
    "reference": "121708",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Sheeba",
    "image": "https://spcadbn.org.za/petimages/_B3sheeba.jpg",
    "description": "Sheeba is a sweet girl with loads of love and affection to share with her human family. She is a 1 year old cross breed. Kennel B3 Ref: 113864.",
    "sex": "female",
    "age": 1.0,
    "kennel": "B3",
    "reference": "113864",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Bonnie",
    "image": "https://spcadbn.org.za/petimages/2024-11-29_A14bonnie.jpg",
    "description": "Bonnie is a sweet girl and playful girl and loves her treats. She is a 1 year old cross breed. Kennel A14 Ref: 121127.",
    "sex": "female",
    "age": 1.0,
    "kennel": "A14",
    "reference": "121127",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Zeus",
    "image": "https://spcadbn.org.za/petimages/_B11zues.jpg",
    "description": "Zeus is an excitable boy and would love a family that will play with him. He is a 1 year old Boerboel cross Ridgeback. Unwanted. Kennel B11 Ref: 121458.",
    "sex": "male",
    "age": 1.0,
    "kennel": "B11",
    "reference": "121458",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Bella",
    "image": "https://spcadbn.org.za/petimages/_A8bella.jpg",
    "description": "Bella is an intelligent girl with a curious nature. She is a 3 year old Labrador cross. Unwanted. Kennel A8 Ref: 118687.",
    "sex": "female",
    "age": 3.0,
    "kennel": "A8",
    "reference": "118687",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Ivy",
    "image": "https://spcadbn.org.za/petimages/_B1ivy.jpg",
    "description": "Ivy is a happy and confident and loves to play with her humans and cuddle with them. She is a 5 year old cross breed. Kennel B1 Ref: 121930.",
    "sex": null,
    "age": 5.0,
    "kennel": "B1",
    "reference": "121930",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Cooper",
    "image": "https://spcadbn.org.za/petimages/_b7cooper.jpg",
    "description": "Cooper is a sweet and gentle little boy with loads of love to give. He is a 3 year old German Shepherd cross. Stray. Kennel B7 Ref: 117663.",
    "sex": "male",
    "age": 3.0,
    "kennel": "B7",
    "reference": "117663",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Bobby",
    "image": "https://spcadbn.org.za/petimages/_B5bobby.jpg",
    "description": "Bobby is an excitable boy with a confident nature. He enjoys going on adventures and being the centre of attention. Bobby is 3 year old Labrador cross. Stray. Kennel B5 Ref: 120769.",
    "sex": "male",
    "age": 3.0,
    "kennel": "B5",
    "reference": "120769",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Layla",
    "image": "https://spcadbn.org.za/petimages/_B2layla.jpg",
    "description": "Layla is a sweet little girl with a gentle nature. She enjoys cuddling and showing her affection. She is a 2 year old German Shepherd cross. Unwanted. Kennel B2 Ref: 120057.",
    "sex": "female",
    "age": 2.0,
    "kennel": "B2",
    "reference": "120057",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Milly",
    "image": "https://spcadbn.org.za/petimages/_a12milly.jpg",
    "description": "Milly is an energetic and playful pooch and loves to show her affection. She is a 2 year old cross breed. Stray. Kennel A12 Ref: 120229.",
    "sex": null,
    "age": 2.0,
    "kennel": "A12",
    "reference": "120229",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Maya",
    "image": "https://spcadbn.org.za/petimages/_a10maya.jpg",
    "description": "Maya is a confident and friendly girl and always fully of doggy smiles. She is a 3 year old Border Collie cross. Unwanted. Kennel A10 Ref: 121067.",
    "sex": "female",
    "age": 3.0,
    "kennel": "A10",
    "reference": "121067",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Queen",
    "image": "https://spcadbn.org.za/petimages/_a5queen.jpg",
    "description": "Queen is a sweet young lady with a calm and well behaved nature. She is a 2 year old cross breed. Unwanted. Kennel A5 Ref: 120240.",
    "sex": null,
    "age": 2.0,
    "kennel": "A5",
    "reference": "120240",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Luna",
    "image": "https://spcadbn.org.za/petimages/_A11luna.jpg",
    "description": "Luna is an excitable and playful girl and loves to bounce around the garden. She is a 3 year old German Shepherd cross. Unwanted. Kennel A11 Ref: 119236.",
    "sex": "female",
    "age": 3.0,
    "kennel": "A11",
    "reference": "119236",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Cruize",
    "image": "https://spcadbn.org.za/petimages/_A6cruze.jpg",
    "description": "Cruize is a gentle giant with a placid nature. He is well behaved and gets on with everyone. Cruize is a 1 year old Boerboel cross. Unwanted. Kennel A6 Ref: 121934.\n",
    "sex": null,
    "age": 1.0,
    "kennel": "A6",
    "reference": "121934",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Roxanne",
    "image": "https://spcadbn.org.za/petimages/_A9Roxanne.jpg",
    "description": "Roxanne is a real little sweetheart with a lovable and playful nature. She is a 7 month old German Shepherd cross. Unwanted. Kennel A9 Ref: 121189.",
    "sex": null,
    "age": 0.5833333333333334,
    "kennel": "A9",
    "reference": "121189",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Marley",
    "image": "https://spcadbn.org.za/petimages/_A2Marley.jpg",
    "description": "Marley has a placid nature and is calm and well behaved. He is a 1 year old Labrador cross. Unwanted. Kennel A2 Ref: 123078.",
    "sex": null,
    "age": 1.0,
    "kennel": "A2",
    "reference": "123078",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Peaches",
    "image": "https://spcadbn.org.za/petimages/_A1peaches.jpg",
    "description": "Peaches has a curious nature and loves to explore her surroundings. She is a 2 year old Labrador cross. Unwanted. Kennel A1 Ref: 121918.",
    "sex": null,
    "age": 2.0,
    "kennel": "A1",
    "reference": "121918",
    "type": "Dog",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Oreo",
    "image": "https://spcadbn.org.za/petimages/_oreo.jpg",
    "description": "Oreo is a handsome boy and loves to keep you company. He is 3 years old. Stray. Ref: 121415.",
    "sex": "male",
    "age": 3.0,
    "kennel": null,
    "reference": "121415",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Tinkerbell",
    "image": "https://spcadbn.org.za/petimages/_tinkerbell.jpg",
    "description": "Tinkerbell is a playful girl with a curious nature. She is 1 year old. Stray. Ref: 117699.\n",
    "sex": "female",
    "age": 1.0,
    "kennel": null,
    "reference": "117699",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Marcy",
    "image": "https://spcadbn.org.za/petimages/_marcy.jpg",
    "description": "Marcy is a playful girl with loads of character. She is 2 years old. Unwanted. Ref: 120843.",
    "sex": "female",
    "age": 2.0,
    "kennel": null,
    "reference": "120843",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Lily",
    "image": "https://spcadbn.org.za/petimages/_lily.jpg",
    "description": "Lily is a sweet girl and loves to have her cat naps high up. She is 1 year old. Stray. Ref:121538.",
    "sex": "female",
    "age": 1.0,
    "kennel": null,
    "reference": "121538",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Kitter-Kat",
    "image": "https://spcadbn.org.za/petimages/_kitter.jpg",
    "description": "Kitter-Kat is a handsome boy with a chilled nature. He is 1 year old. Unwanted. Ref: 121910.",
    "sex": "male",
    "age": 1.0,
    "kennel": null,
    "reference": "121910",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Jetson",
    "image": "https://spcadbn.org.za/petimages/_jetson.jpg",
    "description": " Jestson is a playful boy and loves to follow you around. He is 3 years old. Stray. Ref: 120784.",
    "sex": "male",
    "age": 3.0,
    "kennel": null,
    "reference": "120784",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Jasper",
    "image": "https://spcadbn.org.za/petimages/_jasper.jpg",
    "description": "Jasper is a confident boy and loves to be the centre attention. He is 1 year old. Unwanted. Ref: 123076.",
    "sex": "male",
    "age": 1.0,
    "kennel": null,
    "reference": "123076",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Cleo",
    "image": "https://spcadbn.org.za/petimages/_cleo.jpg",
    "description": "Cleo is a well behaved girl and loves treats and naps. She is an adult. Stray. Ref: 120064.",
    "sex": "female",
    "age": null,
    "kennel": null,
    "reference": "120064",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Callie",
    "image": "https://spcadbn.org.za/petimages/_callie.jpg",
    "description": "Callie is a confident girl and enjoys being in the spotlight. She is 2 years old. Stray. Ref: 121429.",
    "sex": "female",
    "age": 2.0,
    "kennel": null,
    "reference": "121429",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  },
  {
    "name": "Ash",
    "image": "https://spcadbn.org.za/petimages/_Ash.jpg",
    "description": "Ash is a happy boy with a confident and friendly nature. He is an adult. Stray. Ref: 120064.",
    "sex": "male",
    "age": null,
    "kennel": null,
    "reference": "120064",
    "type": "Cat",
    "spca_branch": "Durban and Coast SPCA",
    "province": "KwaZulu-Natal",
    "latitude": -29.80363108540157,
    "longitude": 30.99372385667752
  }
];

const petContainer = document.getElementById('pets');
const searchInput = document.getElementById('search');
const sexSelect = document.getElementById('sex');
const minAgeInput = document.getElementById('minAge');
const maxAgeInput = document.getElementById('maxAge');
const petTypeSelect = document.getElementById('petType');
const searchButton = document.getElementById('searchButton');

function renderPets(filteredPets) {
  petContainer.innerHTML = '';
  filteredPets.forEach(pet => {
    const petCard = document.createElement('div');
    petCard.classList.add('pet-card');
    petCard.innerHTML = `
      <h3>${pet.name}</h3>
      <img src="${pet.image}" alt="${pet.name}" width="300px">
      <p>${pet.description}</p>
    `;
    petContainer.appendChild(petCard);
  });
}

function renderAlertButton() {
  petContainer.innerHTML = '';
  const noPetsCard = document.createElement('div');
  noPetsCard.classList.add('pet-card');
  noPetsCard.innerHTML = `
    <h3>No pets found matching your filters</h3>
    <button disabled>Setup Alert (coming soon)</button>
    <p>We will alert you when a matching pet is added to the system</p>
  `;
  petContainer.appendChild(noPetsCard);
}

function filterPets() {
  const searchTerm = searchInput.value.toLowerCase();
  const selectedSex = sexSelect.value;
  const minAge = parseInt(minAgeInput.value) || 0;
  const maxAge = parseInt(maxAgeInput.value) || Infinity;
  const selectedPetType = petTypeSelect.value;

  const filteredPets = pets.filter(pet => {
    const nameMatches = pet.name.toLowerCase().includes(searchTerm);
    let sexMatches = true;
    if (selectedSex) {
      sexMatches = pet.sex === selectedSex;
    }
    let petAge = pet.age;
    if (petAge === null) {
      petAge = 5; // Treat 'None' as 5 years old
    }
    const ageMatches = petAge >= minAge && petAge <= maxAge;
    let typeMatches = true;
    if (selectedPetType) {
      typeMatches = pet.type === selectedPetType;
    }
    return nameMatches && sexMatches && ageMatches && typeMatches;
  });

  if (filteredPets.length == 0) {
    // We did not find any matching pets, show the setup alert button
    renderAlertButton();
  } else {
    // We found matching pets, render them
    renderPets(filteredPets);
  }
}

searchButton.addEventListener('click', filterPets);
sexSelect.addEventListener('change', filterPets);
minAgeInput.addEventListener('input', filterPets);
maxAgeInput.addEventListener('input', filterPets);
petTypeSelect.addEventListener('change', filterPets);

renderPets(pets); // Render all pets initially