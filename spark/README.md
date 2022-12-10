## Spark-Kafka Integration

This submodule contains the code for the spark job. Currently we are trying to get the spark job to run on the incoming Kafka stream using the code written in ```job.py```. Our findings so far suggest that the logs of the job are stored in the worker nodes under the ```/opt/bitnami/spark/work``` directory. However, these logs are pretty useless since the worker doesn't actually log anything. Furthermore, any files created are actually saved on the master node/the node which runs the job. The only part that actually executes on the spark cluster is anything related to the dataframe. 

Steps to execute the job on the spark cluster:

1. Create the spark using the ```install-spark.sh``` script. This will create a spark cluster in the ```spark-ns``` namespace.
2. Submit ```job.py``` to the cluster for processing:
   1. First open up a bash shell on the master node using:
   ```kubectl exec -n spark-ns -it pod/my-release-spark-master-0 -- bash```
   2. Install the the ```py4j``` package for python using ```pip install py4j```
   3. Copy the job script into master node using:
   ```kubectl cp --namespace spark-ns job.py my-release-spark-master-0:/opt/bitnami/spark/job.py```
   Note that this needs to be done everytime the job script changes. We could just make changes in the container directly but that requires installing nano/vi/vim which isn't possible because the default bash shell given has no user hence, ```apt install``` doesn't work.
   4. Submit the job from inside the master's ```/opt/bitnami/spark/``` directory using the ```spark-submit``` command:
   ```spark-submit --master spark://my-release-spark-master-svc:7077 --packages org.apache.spark:spark-sql-kafka-0-10_2.12:3.3.1 job.py```
   The packages option is necessary to get the job to run.
3. Note that ONLY the dataframe/RDD operations will actually run on the cluster and the master node will execute the rest of the code. This means if the code somehow creates a file, it SHOULD show up in the directory where you launched the code from. 

### Current Revelations:
1. Currently we learnt that we have to use writeStream.start to actually get the dataframe. 
2. We also have to cast the "Key" and "Value" fields of the said data frame to bytes and then parse the JSON string?
3. We have to figure out a way to use the dataframe properly

### Ultimate Objective
The goal of this spark cluster is to analyze the time series data and send a message back to the kafka message board with the topic `notifications` where the message will be of the format `City, State, Product Number`. This will give an alert to the frontend with the required information. Possible extension could be redirecting the product from the nearest available store efficiently.