
import 'package:dio/dio.dart';

class ApiClient {
  final Dio _dio;

  ApiClient(): _dio = Dio(
    BaseOptions(
      baseUrl: const String.fromEnvironment('API_URL', defaultValue: 'http://localhost:8080'),
      connectTimeout: const Duration(seconds: 5000),
      receiveTimeout: const Duration(seconds: 3000),
    )
  );

  Dio get dio => _dio;


}