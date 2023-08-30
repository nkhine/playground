#include <math.h>
#include <stdio.h>

const float theta_spacing = 0.07;
const float phi_spacing   = 0.02;

const float R1 = 1;
const float R2 = 2;
const float K2 = 5;
int screen_width = 80;
int screen_height = 40;

// Calculate K1 based on screen size
const float K1 = screen_width * K2 * 3 / (8 * (R1 + R2));

void render_frame(float A, float B) {
  float cosA = cos(A), sinA = sin(A);
  float cosB = cos(B), sinB = sin(B);

  char output[screen_width][screen_height];
  float zbuffer[screen_width][screen_height];

  // Initialize output and zbuffer
  for (int j = 0; j < screen_height; ++j) {
    for (int i = 0; i < screen_width; ++i) {
      output[i][j] = ' ';
      zbuffer[i][j] = 0;
    }
  }

  for (float theta = 0; theta < 2 * M_PI; theta += theta_spacing) {
    float costheta = cos(theta), sintheta = sin(theta);

    for (float phi = 0; phi < 2 * M_PI; phi += phi_spacing) {
      float cosphi = cos(phi), sinphi = sin(phi);
    
      float circlex = R2 + R1 * costheta;
      float circley = R1 * sintheta;

      float x = circlex * (cosB * cosphi + sinA * sinB * sinphi)
        - circley * cosA * sinB; 
      float y = circlex * (sinB * cosphi - sinA * cosB * sinphi)
        + circley * cosA * cosB;
      float z = K2 + cosA * circlex * sinphi + circley * sinA;
      float ooz = 1 / z;  // "one over z"
      
      int xp = (int) (screen_width / 2 + K1 * ooz * x);
      int yp = (int) (screen_height / 2 - K1 * ooz * y);
      
      float L = cosphi * costheta * sinB - cosA * costheta * sinphi -
        sinA * sintheta + cosB * (cosA * sintheta - costheta * sinA * sinphi);

      if (L > 0) {
        if (xp >= 0 && xp < screen_width && yp >= 0 && yp < screen_height) {
          if (ooz > zbuffer[xp][yp]) {
            zbuffer[xp][yp] = ooz;
            int luminance_index = (int) (L * 8);
            char luminance_chars[12] = ".,-~:;=!*#$@";
            output[xp][yp] = luminance_chars[luminance_index];
          }
        }
      }
    }
  }

  printf("\x1b[H");
  for (int j = 0; j < screen_height; ++j) {
    for (int i = 0; i < screen_width; ++i) {
      putchar(output[i][j]);
    }
    putchar('\n');
  }
}

