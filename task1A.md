# Solution Design

## Requirements

Please briefly explain how you would design a data ingestion and storage solution,
outlining the various components (eg. database), how/what data flows between
them, and what considerations made you choose this design/architecture.
If you have experience with Microsoft Azure, please refer to specific Azure services
that your solution might rely upon.

## Solution

* Table in DB (SQL) storing meter serial ID, date of reading, and reading
* Transmitted data added to stremm (kafka or similar) to be queued for processing - meter ID, timestamp added to queue, reading - data can be ingested when resources are available, available by different users
* Serverless process to process batches, data standardisation if required - multiple can be span up in parallel based on requirements
* Monitoring (Kibana or similar) to ensure everything behaving as expected