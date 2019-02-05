# Design

## Use Cases

- Scrape data (2 stages)
    - scrape
    - format data from scraped
  
- list spells
- list equipment
- list traits
- list creatures
- filter
- flag short/long output
- describe
- full-text search
- generate store
- generate engagement
- character inventory
- character spells(known, prepared)
- character stats
- create new items, spells, traits, rules
- generate npc

## Getting Started

	$ git clone https://github.com/njdaniel/dnd.git
	$ go install 
	
	
## Designs

### Design #1

Decentralized.

Pull all data into directories and json files. Search with basic linux tools

### Design #2 

Centralized. 

REST API/gRPC?

Good for querying for data. 
Adding new features, deployments.
Goes with k8s.

How would cli work? AWS s3 cli still uses the linux directory approach. Personally I would like something like this.

From scratch/framework(buffalo). Code generation is quite nice.